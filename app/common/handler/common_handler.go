package handler

import (
	"encoding/json"
	commondto "gitlab.dot.co.id/playground/boilerplates/golang-service/app/common/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/common/middleware"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CommonHandler struct {
}

func NewCommonHandler(router *httprouter.Router) {
	basePath := "/api/v1/commons"
	projectHandler := &CommonHandler{}

	router.GET(basePath+"/our-service", middleware.Wrapper(projectHandler.OurServiceList(), middleware.MiddlewareType{XServiceAuthToken: true}))
	router.GET(basePath+"/target-time", middleware.Wrapper(projectHandler.TargetTimeList(), middleware.MiddlewareType{XServiceAuthToken: true}))
}

// @Summary Our Service
// @Description Our Service
// @Tags Commons
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param basic-param query commondto.Pagination true "basic param"
// @Success 200 {object} pkg.BaseResponse{data=[]commondto.ListServiceResponse} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/commons/our-service [get]
func (h *CommonHandler) OurServiceList() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		res := []commondto.ListServiceResponse{
			{
				Name:  "Web Aplikasi",
				Value: "WEB_APPLICATION",
			},
			{
				Name:  "Website Compro",
				Value: "WEB_COMPRO",
			},
			{
				Name:  "Mobile Apps",
				Value: "MOBILE_APPS",
			},
			{
				Name:  "Maintenance Aplikasi",
				Value: "MAINTENANCE",
			},
		}

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data:    res,
			Meta: &pkg.MetaResponse{
				Page:    1,
				PerPage: 10,
			},
		})

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(dataByte)
		if err != nil {
			return
		}
	}
}

// @Summary Target Time
// @Description Target Time
// @Tags Commons
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param basic-param query commondto.Pagination true "basic param"
// @Success 200 {object} pkg.BaseResponse{data=[]commondto.ListTargetTimeResponse} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/commons/target-time [get]
func (h *CommonHandler) TargetTimeList() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		res := []commondto.ListTargetTimeResponse{
			{
				Name:  "Sangat Cepat 1 Minggu",
				Value: "ONE_WEEKS",
			},
			{
				Name:  "Cukup Cepat - 1 Bulan",
				Value: "ONE_MONTH",
			},
			{
				Name:  "Cukup Lama 1-3 Bulan",
				Value: "THREE_MONTH",
			},
			{
				Name:  "Jangka Panjang",
				Value: "LONG_TERM",
			},
		}

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data:    res,
			Meta: &pkg.MetaResponse{
				Page:    1,
				PerPage: 10,
			},
		})

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(dataByte)
		if err != nil {
			return
		}
	}
}
