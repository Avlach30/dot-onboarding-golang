package bannerdto

type BannerListReq struct {
	Page    int `json:"page" query:"page"`
	PerPage int `json:"per_page" query:"per_page"`
}
