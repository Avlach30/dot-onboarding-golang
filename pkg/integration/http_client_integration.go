package integration

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	encodedUrl "net/url"
	"strconv"
	"strings"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
)

func (client *Client) Get(endpoint string, headers *Headers, responseBody any) (any, error) {
	response, err := client.SendHTTPRequest(http.MethodGet, endpoint, headers, nil, responseBody)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client *Client) Post(endpoint string, headers *Headers, requestBody interface{}, responseBody any) (any, error) {
	response, err := client.SendHTTPRequest(http.MethodPost, endpoint, headers, requestBody, responseBody)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (client *Client) Put(endpoint string, headers *Headers, requestBody interface{}, responseBody any) (any, error) {
	response, err := client.SendHTTPRequest(http.MethodPut, endpoint, headers, requestBody, responseBody)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (client *Client) Patch(endpoint string, headers *Headers, requestBody interface{}, responseBody any) (any, error) {
	response, err := client.SendHTTPRequest(http.MethodPatch, endpoint, headers, requestBody, responseBody)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (client *Client) Delete(endpoint string, headers *Headers, responseBody any) (any, error) {
	response, err := client.SendHTTPRequest(http.MethodDelete, endpoint, headers, nil, responseBody)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (client *Client) SendHTTPRequest(method, endpoint string, headers *Headers, requestBody interface{}, responseBody any) (any, error) {
	url := client.BaseURL + endpoint
	externalCircuitBreaker := singleton.GetCircuitBreaker(singleton.ExternalCircuitBreaker)

	isCircuitBreakerEnable, err := strconv.ParseBool(config.IsCircuitBreakerEnabled)

	isReadyToTrip := externalCircuitBreaker.IsReadyToTrip()
	if !isReadyToTrip && isCircuitBreakerEnable {
		return nil, fmt.Errorf("Error circuit breaker is open, cannot send request to %s", url)
	}

	reqBody, err := prepareRequestBody(headers.ContentType, requestBody)
	if err != nil {
		log.Printf("Error preparing request body: %v", err)
		return nil, err
	}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	setHeaders(req, headers)
	log.Println("=========================START HTTP REQUEST=========================")
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}
	createLogIntegration(req, resp)
	log.Println("==========================END HTTP REQUEST==========================")

	if isCircuitBreakerEnable {
		externalCircuitBreaker.CountRequest()
		if resp.StatusCode >= http.StatusInternalServerError {
			externalCircuitBreaker.FailureHappend(endpoint)
		}
	}

	if resp.StatusCode >= http.StatusInternalServerError {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&responseBody)

	if err != nil {
		log.Printf("Error decoding response: %v", err)
		return nil, err
	}

	return responseBody, nil
}

func prepareRequestBody(contentType string, requestBody interface{}) (io.Reader, error) {
	switch contentType {
	case constant.ContentTypeJSON:
		return jsonBody(requestBody)
	case constant.ContentTypeXML:
		return xmlBody(requestBody)
	case constant.ContentTypeMultipartFormData:
		return multipartFormBody(requestBody)
	case constant.ContentTypeForm:
		return formBody(requestBody)
	default:
		return nil, fmt.Errorf("unsupported content type: %s", contentType)
	}
}

func jsonBody(requestBody interface{}) (io.Reader, error) {
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(jsonBody), nil
}

func xmlBody(requestBody interface{}) (io.Reader, error) {
	xmlBody, err := xml.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(xmlBody), nil
}

func multipartFormBody(requestBody interface{}) (io.Reader, error) {
	formBody := &bytes.Buffer{}
	writer := multipart.NewWriter(formBody)
	for key, value := range requestBody.(map[string]string) {
		err := writer.WriteField(key, value)
		if err != nil {
			return nil, err
		}
	}
	err := writer.Close()
	if err != nil {
		return nil, err
	}
	return formBody, nil
}

func formBody(requestBody interface{}) (io.Reader, error) {
	formData := encodedUrl.Values{}
	for key, value := range utils.StructToMap(requestBody, true) {
		formData.Add(key, value)
	}
	return strings.NewReader(formData.Encode()), nil
}

func setHeaders(req *http.Request, headers *Headers) {
	req.Header.Set("Content-Type", headers.ContentType)
	req.Header.Set("Authorization", headers.Authorization)
}
