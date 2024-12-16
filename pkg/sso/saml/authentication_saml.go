package saml

import (
	"log"

	saml "github.com/mattbaird/gosaml"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
)

func GenerateSAML() (string, error) {
	log.Println("=========================START GENERATE SAML=========================")

	samlAssertationServiceUrl := config.SAMLAssertionServiceUrl
	samlIssuer := config.SAMLIssuer
	samlCertPath := config.SAMLCertPath
	samlKeyPath := config.SAMLKeyPath

	log.Printf("SAML Configuration: AssertionServiceUrl=%s, Issuer=%s, CertPath=%s, KeyPath=%s",
		samlAssertationServiceUrl, samlIssuer, samlCertPath, samlKeyPath)

	// Configure the app and account settings
	log.Println("Configuring app settings")
	appSettings := saml.NewAppSettings(samlAssertationServiceUrl, samlIssuer)
	log.Println("Configuring account settings")
	accountSettings := saml.NewAccountSettings("cert", samlAssertationServiceUrl)

	// Construct an AuthnRequest
	log.Println("Constructing AuthnRequest")
	authRequest := saml.NewAuthorizationRequest(*appSettings, *accountSettings)

	// Return a SAML AuthnRequest as a string
	log.Println("Getting signed request")
	samlStr, err := authRequest.GetSignedRequest(false, samlCertPath, samlKeyPath)

	if err != nil {
		log.Printf("Error generating SAML : %v", err)
		return "", err
	}

	log.Println("==========================END GENERATE SAML==========================")
	return samlStr, nil
}
