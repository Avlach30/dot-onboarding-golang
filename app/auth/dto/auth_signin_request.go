package dto

type AuthSignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthSignLDAPRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthSignOIDCRequest struct {
	Code string `json:"code" binding:"required"`
}
