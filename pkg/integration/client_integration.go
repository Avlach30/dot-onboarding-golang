package integration

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

type Headers struct {
	Authorization string
	ContentType   string
	SOAPAction    string
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

func createLogIntegration(request *http.Request, response *http.Response) error {

	logRequest(request)
	logResponse(response)

	// do whatever u want to do with log, insert DB, or etc

	return nil
}

func logResponse(response *http.Response) string {
	var responseBody []byte

	log.Printf("Response Status : %s", response.Status)
	log.Printf("Response Headers : %v", response.Header)

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		err := fmt.Sprintf("Error reading response body : %s", err.Error())
		log.Println(err)
		return err
	}

	response.Body = io.NopCloser(bytes.NewBuffer(responseBody))

	log.Printf("Response Body : %s", string(responseBody))

	return string(responseBody)
}

func logRequest(request *http.Request) string {
	var requestBody []byte

	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		err := fmt.Sprintf("Error reading request body : %v", err.Error())
		log.Println(err)
		return err
	}

	request.Body = io.NopCloser(bytes.NewBuffer(requestBody))

	log.Printf("Request %s URL : %s", request.Method, request.URL)
	log.Printf("Request Headers : %s", request.Header)
	log.Printf("Request Body : %s", string(requestBody))

	return string(requestBody)
}
