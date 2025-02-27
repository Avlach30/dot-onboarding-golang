package integration

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
)

func (client *Client) SendSOAPRequest(action string, headers *Headers, requestBody interface{}, responseBody any) (any, error) {
	var resp *http.Response

	externalCircuitBreaker := singleton.GetCircuitBreaker(singleton.ExternalCircuitBreaker)
	isCircuitBreakerEnable, _ := strconv.ParseBool(config.IsCircuitBreakerExternalEnabled)
	isReadyToTrip := externalCircuitBreaker.IsReadyToTrip()
	if !isReadyToTrip && isCircuitBreakerEnable {
		return nil, fmt.Errorf("Error circuit breaker is open, cannot send action to %s", action)
	} else if isCircuitBreakerEnable {
		defer func() {
			externalCircuitBreaker.CountRequest()
			if resp.StatusCode >= http.StatusInternalServerError {
				externalCircuitBreaker.FailureHappend(action)
			}
		}()
	}

	log.Println("========================= START SOAP REQUEST =========================")
	envelope := createSOAPEnvelope(requestBody)
	xmlBody, err := marshalEnvelope(envelope)
	log.Println("========================== END SOAP REQUEST ==========================")

	if err != nil {
		return nil, err
	}

	req, err := createHTTPRequest(client.BaseURL, xmlBody, action, headers)
	if err != nil {
		return nil, err
	}

	resp, err = sendRequest(client.HTTPClient, req)
	if err != nil {
		return nil, err
	}

	createLogIntegration(req, resp)

	defer resp.Body.Close()

	return processResponse(resp, responseBody)
}

func createSOAPEnvelope(requestBody interface{}) interface{} {
	return struct {
		XMLName xml.Name `xml:"soap:Envelope"`
		Body    struct {
			Content interface{} `xml:",any"`
		} `xml:"soap:Body"`
	}{
		Body: struct {
			Content interface{} `xml:",any"`
		}{Content: requestBody},
	}
}

func marshalEnvelope(envelope interface{}) ([]byte, error) {
	xmlBody, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		log.Printf("Error marshaling SOAP envelope: %v", err)
		return nil, err
	}
	return []byte(xml.Header + string(xmlBody)), nil
}

func createHTTPRequest(url string, xmlBody []byte, action string, headers *Headers) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(xmlBody))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", action)
	if headers.Authorization != "" {
		req.Header.Set("Authorization", headers.Authorization)
	}

	return req, nil
}

func sendRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	log.Println("========================= START SOAP REQUEST =========================")
	logRequest(req)
	resp, err := client.Do(req)
	logResponse(resp)
	log.Println("========================== END SOAP REQUEST ==========================")

	if err != nil {
		log.Printf("Error sending request: %v", err)
		return resp, err
	}

	return resp, nil
}

func processResponse(resp *http.Response, responseBody any) (any, error) {
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, err
	}

	err = xml.Unmarshal(respBody, &responseBody)
	if err != nil {
		log.Printf("Error unmarshaling SOAP response: %v", err)
		return nil, err
	}

	return responseBody, nil
}
