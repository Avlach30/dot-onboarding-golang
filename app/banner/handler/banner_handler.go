package handler

import (
	"encoding/json"
	"github.com/codespace-id/codespace-x/app/banner/bannerdomain"
	"github.com/codespace-id/codespace-x/app/banner/bannerdto"
	httperror "github.com/codespace-id/codespace-x/pkg/common/error"
	"github.com/codespace-id/codespace-x/pkg/common/middleware"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/codespace-id/codespace-x/pkg"
	"github.com/julienschmidt/httprouter"
)

type BannerHandler struct {
	bannerUsecase bannerdomain.Usecase
}

func NewBannerHandler(router *httprouter.Router, bannerUsecase bannerdomain.Usecase) {
	basePath := "/api/v1/banners"
	bannerHandler := &BannerHandler{
		bannerUsecase: bannerUsecase,
	}

	router.GET(basePath, middleware.Wrapper(bannerHandler.ListBanner(), middleware.MiddlewareType{TokenAuth: true, XServiceAuthToken: true}))

}

// @Summary List Banner
// @Description List Banner
// @Tags Banner
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param authorization header string true "Authorization value"
// @Param basic-param query bannerdto.BannerListReq true "basic param"
// @Success 200 {object} pkg.BaseResponse{data=[]bannerdto.BannerResponse} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/banners [get]
func (h *BannerHandler) ListBanner() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		var err error
		var payloadReq bannerdto.BannerListReq

		queryParams := r.URL.Query()

		if page, ok := queryParams["page"]; ok {
			pageAsInt, _ := strconv.Atoi(page[0])
			payloadReq.Page = pageAsInt
		}
		if perPage, ok := queryParams["per_page"]; ok {
			perPageInt, _ := strconv.Atoi(perPage[0])
			payloadReq.PerPage = perPageInt
		}

		if payloadReq.Page == 0 {
			payloadReq.Page = 1
		}
		if payloadReq.PerPage == 0 {
			payloadReq.PerPage = 10
		}

		// validate payload
		errMsgs := pkg.ValidateStruct(payloadReq)
		if len(errMsgs) > 0 {
			httperror.SetResponse(w, 400, errMsgs)
			return
		}
		defer r.Body.Close()

		res, err := h.bannerUsecase.Get(r.Context(), payloadReq.Page, payloadReq.PerPage)
		if err != nil {
			log.Println("error getting banners: ", string(debug.Stack()))
			httperror.SetResponse(w, 500, "internal server error")
			return
		}

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data:    res,
			Meta: &pkg.MetaResponse{
				Page:    payloadReq.Page,
				PerPage: payloadReq.PerPage,
			},
		})

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(dataByte)
		if err != nil {
			return
		}
	}
}
