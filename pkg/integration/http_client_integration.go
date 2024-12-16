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
	"strings"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

type Headers struct {
	Authorization string
	ContentType   string
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

func (client *Client) SendRequest(method, endpoint string, headers *Headers, requestBody interface{}, responseBody any) (any, error) {
	url := client.BaseURL + endpoint
	var reqBody io.Reader

	switch headers.ContentType {
	case constant.ContentTypeJSON:
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			log.Printf("Error marshaling JSON: %v", err)
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonBody)
	case constant.ContentTypeXML:
		xmlBody, err := xml.Marshal(requestBody)
		if err != nil {
			log.Printf("Error marshaling XML: %v", err)
			return nil, err
		}
		reqBody = bytes.NewBuffer(xmlBody)
	case constant.ContentTypeMultipartFormData:
		formBody := &bytes.Buffer{}
		writer := multipart.NewWriter(formBody)
		for key, value := range requestBody.(map[string]string) {
			err := writer.WriteField(key, value)
			if err != nil {
				log.Printf("Error writing form field: %v", err)
				return nil, err
			}
		}
		err := writer.Close()
		if err != nil {
			log.Printf("Error closing multipart writer: %v", err)
			return nil, err
		}
		reqBody = formBody
		headers.ContentType = writer.FormDataContentType()
	case constant.ContentTypeForm:
		formData := encodedUrl.Values{}
		for key, value := range utils.StructToMap(requestBody, true) {
			formData.Add(key, value)
		}
		reqBody = strings.NewReader(formData.Encode())
	default:
		log.Printf("Unsupported content type: %s", headers.ContentType)
		return responseBody, fmt.Errorf("unsupported content type: %s", headers.ContentType)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", headers.ContentType)
	req.Header.Set("Authorization", headers.Authorization)

	log.Println("=========================START HTTP REQUEST=========================")
	logRequest(req)
	resp, err := client.HTTPClient.Do(req)
	logResponse(resp)
	log.Println("==========================END HTTP REQUEST==========================")

	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	resBodyReader := resp.Body
	json.NewDecoder(resBodyReader).Decode(&responseBody)

	return responseBody, nil
}

func (client *Client) Get(endpoint string, headers *Headers, responseBody any) (any, error) {
	return client.SendRequest(http.MethodGet, endpoint, headers, nil, responseBody)
}

func (client *Client) Post(endpoint string, headers *Headers, requestBody interface{}, responseBody any) (any, error) {
	return client.SendRequest(http.MethodPost, endpoint, headers, requestBody, responseBody)
}

func (client *Client) Put(endpoint string, headers *Headers, requestBody interface{}, responseBody any) (any, error) {
	return client.SendRequest(http.MethodPut, endpoint, headers, requestBody, responseBody)
}

func (client *Client) Patch(endpoint string, headers *Headers, requestBody interface{}, responseBody any) (any, error) {
	return client.SendRequest(http.MethodPatch, endpoint, headers, requestBody, responseBody)
}

func (client *Client) Delete(endpoint string, headers *Headers, responseBody any) (any, error) {
	return client.SendRequest(http.MethodDelete, endpoint, headers, nil, responseBody)
}

func logResponse(response *http.Response) {
	var responseBody []byte

	log.Printf("Response Status : %s", response.Status)
	log.Printf("Response Headers : %v", response.Header)

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response body : %v", err)
		return
	}

	defer response.Body.Close()

	log.Printf("Response Body : %s", string(responseBody))
}

func logRequest(request *http.Request) {
	var requestBody []byte

	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		return
	}

	defer request.Body.Close()

	log.Printf("Request %s URL : %s", request.Method, request.URL)
	log.Printf("Request Headers : %s", request.Header)
	log.Printf("Request Body : %s", utils.StructToMap(requestBody, true))
}
