package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codespace-id/codespace-x/app/dto"
	"github.com/codespace-id/codespace-x/pkg/common/middleware"

	"github.com/codespace-id/codespace-x/pkg"
	"github.com/julienschmidt/httprouter"
)

type BannerHandler struct {
}

func NewBannerHandler(router *httprouter.Router) {
	basePath := "/api/v1/banners"
	bannerHandler := &BannerHandler{}

	router.GET(basePath, middleware.Wrapper(bannerHandler.ListBanner(), middleware.MiddlewareType{TokenAuth: true, XServiceAuthToken: true}))

}

// @Summary List Banner
// @Description List Banner
// @Tags Banner
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param authorization header string true "Authorization value"
// @Success 200 {object} pkg.BaseResponse{data=[]dto.BannerResponse} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/banners [get]
func (h *BannerHandler) ListBanner() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data: []dto.BannerResponse{
				{
					Name:        "Promo Wordpress Landing",
					Description: "Dapatkan wordpress landing",
					SourceURL:   "https://res.cloudinary.com/deafomwc7/image/upload/v1664837512/codespace/images/portfolio/Doyan_zae25x.png",
				},
			},
		})

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(dataByte)
		if err != nil {
			return
		}
	}
}
