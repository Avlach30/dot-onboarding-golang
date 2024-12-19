package integration

import (
	"bytes"
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

func NewHTTPClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

func (client *Client) Get(endpoint string, headers *Headers, responseBody any) (any, error) {
	return client.SendHTTPRequest(http.MethodGet, endpoint, headers, nil, responseBody)
}

func (client *Client) Post(endpoint string, headers *Headers, requestBody interface{}, responseBody any) (any, error) {
	return client.SendHTTPRequest(http.MethodPost, endpoint, headers, requestBody, responseBody)
}

func (client *Client) Put(endpoint string, headers *Headers, requestBody interface{}, responseBody any) (any, error) {
	return client.SendHTTPRequest(http.MethodPut, endpoint, headers, requestBody, responseBody)
}

func (client *Client) Patch(endpoint string, headers *Headers, requestBody interface{}, responseBody any) (any, error) {
	return client.SendHTTPRequest(http.MethodPatch, endpoint, headers, requestBody, responseBody)
}

func (client *Client) Delete(endpoint string, headers *Headers, responseBody any) (any, error) {
	return client.SendHTTPRequest(http.MethodDelete, endpoint, headers, nil, responseBody)
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

	response.Body = io.NopCloser(bytes.NewBuffer(responseBody))

	log.Printf("Response Body : %s", string(responseBody))
}

func logRequest(request *http.Request) {
	var requestBody []byte

	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		return
	}

	request.Body = io.NopCloser(bytes.NewBuffer(requestBody))

	log.Printf("Request %s URL : %s", request.Method, request.URL)
	log.Printf("Request Headers : %s", request.Header)
	log.Printf("Request Body : %s", string(requestBody))
}
