package integration

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/integration/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/integration/repository"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
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

	req := logRequest(request)
	res := logResponse(response)

	if config.LogDriver != "database" {
		return nil
	}

	logIntegrationRepo := repository.NewLogIntegrationRepository(singleton.GetDBUtil())
	return logIntegrationRepo.CreateLogIntegration(&domain.LogIntegrationEntity{
		URL:      request.Method + " " + request.URL.String(),
		Request:  req,
		Response: res,
		Status:   response.Status,
		Scheme:   "HTTP",
	})
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
