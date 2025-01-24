package oidc

import (
	"log"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/integration"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/sso/oidc/dto"
)

func GetAuthToken(code string, redirectUri string) (string, error) {
	log.Println("=========================START GET AUTH TOKEN OIDC=========================")

	grantType := config.OIDCGrantType
	clientId := config.OIDCClientId
	clientSecret := config.OIDCClientSecret
	tokenUrl := config.OIDCTokenUrl

	formData := &dto.GetTokenOIDCRequest{
		GrantType:    grantType,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Code:         code,
		RedirectUri:  redirectUri,
	}

	httpClientGetToken := integration.NewClient("")

	bearerTokenSSO := ""
	{
		responseBody, err := httpClientGetToken.Post(tokenUrl, &integration.Headers{
			ContentType: constant.ContentTypeForm,
		}, formData, &dto.GetTokenOIDCResponse{})
		if err != nil {
			log.Printf("Error in GetAuthToken : %v", err)
			return bearerTokenSSO, err
		}

		tokenResponse := responseBody.(*dto.GetTokenOIDCResponse)
		bearerTokenSSO = tokenResponse.TokenType + " " + tokenResponse.AccessToken
	}

	log.Println("==========================END GET AUTH TOKEN OIDC==========================")
	return bearerTokenSSO, nil
}

func GetEmail(bearerTokenSSO string) (string, error) {
	log.Println("==========================START GET EMAIL OIDC==========================")

	userInfoUrl := config.OIDCUserInfoUrl

	emailSSO := ""
	{
		headersUserInfo := integration.Headers{
			Authorization: bearerTokenSSO,
			ContentType:   constant.ContentTypeJSON,
		}

		httpClientGetToken := integration.NewClient("")
		responseBody, err := httpClientGetToken.Get(userInfoUrl, &headersUserInfo, &dto.GetUserInfoResponse{})
		userInfoResponse := responseBody.(*dto.GetUserInfoResponse)

		if err != nil || !userInfoResponse.EmailVerified {
			log.Printf("Error in GetEmail : %v", err)
			return "", err
		}

		userInfoResponse.EmailVerified = true
		emailSSO = userInfoResponse.Email
	}

	log.Println("==========================END GET EMAIL OIDC==========================")
	return emailSSO, nil
}

func GetEmailByCode(code string, redirectUri string) (string, error) {
	log.Println("=========================START GET EMAIL BY CODE OIDC=========================")

	emailSSO := ""
	bearerTokenSSO, err := GetAuthToken(code, redirectUri)
	if err != nil {
		log.Printf("Error in GetEmailByCode : %v", err)
		return emailSSO, err
	}

	email, err := GetEmail(bearerTokenSSO)
	if err != nil {
		log.Printf("Error in GetEmailByCode : %v", err)
		return emailSSO, err
	}

	log.Println("=========================END GET EMAIL BY CODE OIDC=========================")
	return email, nil
}
