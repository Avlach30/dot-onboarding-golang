package config

var (
	Secret           = GetRequired("JWT_SECRET")
	JwtExpiredInDays = GetRequired("JWT_EXPIRED_IN_DAYS")
	ServiceAuthToken = GetRequired("SERVICE_AUTH_TOKEN")

	OIDCBaseUrl             = Get("OIDC_BASE_URL", "")
	OIDCClientId            = Get("OIDC_CLIENT_ID", "")
	OIDCClientSecret        = Get("OIDC_CLIENT_SECRET", "")
	OIDCScope               = Get("OIDC_SCOPE", "")
	OIDCRedirectCallbackUrl = Get("OIDC_REDIRECT_CALLBACK_URL", "")
	OIDCResponseType        = Get("OIDC_RESPONSE_TYPE", "")
	OIDCAuthorizationUrl    = Get("OIDC_AUTHORIZATION_URL", "")
	OIDCUserInfoUrl         = Get("OIDC_USERINFO_URL", "")
	OIDCTokenUrl            = Get("OIDC_TOKEN_URL", "")
	OIDCGrantType           = Get("OIDC_GRANT_TYPE", "")

	SAMLAssertionServiceUrl = Get("SAML_ASSERTION_SERVICE_URL", "")
	SAMLIssuer              = Get("SAML_ISSUER", "")
	SAMLCertPath            = Get("SAML_CERT_PATH", "")
	SAMLKeyPath             = Get("SAML_KEY_PATH", "")
	SAMLNameIdFormat        = Get("SAML_NAMEID_FORMAT", "")
	SAMLSpEntityId          = Get("SAML_SP_ENTITY_ID", "")

	LDAPServer   = Get("LDAP_SERVER", "")
	LDAPPort     = Get("LDAP_PORT", "")
	LDAPBindDN   = Get("LDAP_BINDDN", "")
	LDAPPassword = Get("LDAP_PASSWORD", "")
	LDAPSearchDN = Get("LDAP_SEARCHDN", "")
)
