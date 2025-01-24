package dto

type GetTokenOIDCRequest struct {
	GrantType    string `form:"grant_type"`
	ClientId     string `form:"client_id"`
	ClientSecret string `form:"client_secret"`
	Code         string `form:"code"`
	RedirectUri  string `form:"redirect_uri"`
}
