package ldap

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/sso/ldap/dto"
)

func AuthUsingLDAP(username, password string) (*dto.LDAPUserData, error) {
	log.Println("=========================START GENERATE SAML=========================")

	ldapServer := config.LDAPServer
	ldapPort := config.LDAPPort
	ldapBindDN := config.LDAPBindDN
	ldapPassword := config.LDAPPassword
	ldapSearchDN := config.LDAPSearchDN

	ldapUrl := fmt.Sprintf("ldap://%s:%s", ldapServer, ldapPort)
	log.Printf("Connecting to LDAP server : %s", ldapUrl)
	l, err := ldap.DialURL(ldapUrl)
	if err != nil {
		log.Printf("Error connecting to LDAP server : %v", err)
		return nil, err
	}

	log.Println("Binding to LDAP server")
	err = l.Bind(ldapBindDN, ldapPassword)
	if err != nil {
		log.Printf("Error binding to LDAP server : %v", err)
		return nil, err
	}

	defer l.Close()

	log.Printf("Searching for user : %s", username)
	searchRequest := ldap.NewSearchRequest(
		ldapSearchDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", username),
		[]string{"dn", "cn", "sn", "mail"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Printf("Error searching for user : %v", err)
		return nil, err
	}

	if len(sr.Entries) == 0 {
		log.Printf("User not found : %s", username)
		err := fmt.Errorf("User not found")
		return nil, err
	}

	entry := sr.Entries[0]
	log.Println("Binding as user to verify password")
	err = l.Bind(entry.DN, password)
	if err != nil {
		log.Printf("Error binding as user : %v", err)
		return nil, err
	}

	log.Println("Extracting user data")
	userLDAPData := &dto.LDAPUserData{}
	for _, attr := range entry.Attributes {
		switch attr.Name {
		case "sn":
			userLDAPData.Name = attr.Values[0]
		case "mail":
			userLDAPData.Email = attr.Values[0]
		case "cn":
			userLDAPData.FullName = attr.Values[0]
		}
	}

	log.Println("==========================END GENERATE SAML==========================")
	return userLDAPData, nil
}
