package otp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/Integrations/otp/dto"
)

func CallAPI(httpMethod, url string, bodyPayload []byte) (apiRes *dto.OtpResponse, err error) {
	var apiReq *http.Request

	if bodyPayload != nil {
		apiReq, err = http.NewRequest(httpMethod, url, bytes.NewBuffer(bodyPayload))
	} else {
		apiReq, err = http.NewRequest(httpMethod, url, nil)
	}
	if err != nil {
		return nil, err
	}

	apiReq.Header.Add("content-type", "application/json")
	apiReq.Header.Add("accept", "application/json")

	resp, err := http.DefaultClient.Do(apiReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var otpResponse dto.OtpResponse
	if err = json.Unmarshal(body, &otpResponse); err != nil {
		return nil, err
	}

	return &otpResponse, nil
}
