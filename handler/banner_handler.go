package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codespace-id/codespace-x/pkg"
	"github.com/julienschmidt/httprouter"
)

type BannerHandler struct {
}

func NewBannerHandler(router *httprouter.Router) {
	basePath := "/api/v1/banners"
	bannerHandler := &BannerHandler{}

	router.GET(basePath, bannerHandler.ListBanner())

}

func (h *BannerHandler) ListBanner() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		data := []map[string]interface{}{
			{
				"name":        "Promo Wordpress Landing",
				"description": "Dapatkan wordpress landing",
				"source_url":  "https://res.cloudinary.com/deafomwc7/image/upload/v1664837512/codespace/images/portfolio/Doyan_zae25x.png",
			},
		}

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data:    data,
		})

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(dataByte)
		if err != nil {
			return
		}
	}
}
