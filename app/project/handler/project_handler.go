package handler

import (
	"encoding/json"
	commondto "github.com/codespace-id/codespace-x/app/common/dto"
	projectDomain "github.com/codespace-id/codespace-x/app/project/domain"
	"github.com/codespace-id/codespace-x/app/project/dto"
	userdto "github.com/codespace-id/codespace-x/app/user/dto"
	"github.com/codespace-id/codespace-x/pkg"
	httperror "github.com/codespace-id/codespace-x/pkg/common/error"
	"github.com/codespace-id/codespace-x/pkg/common/middleware"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
)

type ProjectHandler struct {
	projectUsecase       projectDomain.Usecase
	projectPublicUsecase projectDomain.PublicUsecase
}

func NewProjectHandler(router *httprouter.Router, projectUsecase projectDomain.Usecase, projectPublicUsecase projectDomain.PublicUsecase) {
	basePath := "/api/v1/projects"
	projectHandler := &ProjectHandler{
		projectUsecase:       projectUsecase,
		projectPublicUsecase: projectPublicUsecase,
	}

	router.GET(basePath, middleware.Wrapper(projectHandler.ListProject(), middleware.MiddlewareType{TokenAuth: true, XServiceAuthToken: true}))
	router.GET(basePath+"/:uuid", middleware.Wrapper(projectHandler.DetailProject(), middleware.MiddlewareType{TokenAuth: true, XServiceAuthToken: true}))
	router.POST(basePath, middleware.Wrapper(projectHandler.CreateProject(), middleware.MiddlewareType{TokenAuth: true, XServiceAuthToken: true}))
	router.PATCH(basePath+"/:uuid", middleware.Wrapper(projectHandler.UpdateProject(), middleware.MiddlewareType{TokenAuth: true, XServiceAuthToken: true}))
	router.GET(basePath+"/:uuid/histories", middleware.Wrapper(projectHandler.ProjectHistories(), middleware.MiddlewareType{TokenAuth: true, XServiceAuthToken: true}))

}

// @Summary List Project
// @Description List Project
// @Tags Projects
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param authorization header string false "Authorization value"
// @Param basic-param query commondto.Pagination true "basic param"
// @Success 200 {object} pkg.BaseResponse{data=[]dto.ListProjectResponse} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/projects [get]
func (h *ProjectHandler) ListProject() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		// Retrieve values from context (locals)
		phoneNumber, _ := r.Context().Value(middleware.PhoneNumber).(string)
		roles, _ := r.Context().Value(middleware.Roles).(string)

		var err error
		var payloadReq commondto.Pagination

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

		var res []dto.ListProjectResponse

		if phoneNumber != "" {
			res, err = h.projectUsecase.ListProject(r.Context(), phoneNumber, roles, payloadReq.Page, payloadReq.PerPage)
			if err != nil {
				log.Println("error getting projects: ", string(debug.Stack()))
				httperror.SetResponse(w, 500, "internal server error")
				return
			}
		} else if phoneNumber == "" {
			res, err = h.projectPublicUsecase.ListProject(r.Context(), phoneNumber, roles, payloadReq.Page, payloadReq.PerPage)
			if err != nil {
				log.Println("error getting projects: ", string(debug.Stack()))
				httperror.SetResponse(w, 500, "internal server error")
				return
			}
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

// @Summary Detail Project
// @Description Detail Project
// @Tags Projects
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param authorization header string false "Authorization value"
// @Param project_uuid path string true "project_uuid"
// @Success 200 {object} pkg.BaseResponse{data=dto.ProjectDetailResponse} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/projects/{project_uuid} [get]
func (h *ProjectHandler) DetailProject() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		// Retrieve values from context (locals)
		phoneNumber, _ := r.Context().Value(middleware.PhoneNumber).(string)
		UUID := ps.ByName("uuid")

		data, err := h.projectUsecase.ProjectDetail(r.Context(), UUID)
		if err != nil {
			log.Println("error getting project: ", string(debug.Stack()))
			httperror.SetResponse(w, 500, "internal server error")
			return
		}

		if phoneNumber == "" {
			data.Astrodevs = make([]userdto.GetProfileTalentResponse, 0)
		}

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data:    data,
		})

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(dataByte)
		if err != nil {
			return
		}
	}
}

// @Summary CreateTx Project
// @Description CreateTx Project
// @Tags Projects
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param authorization header string true "Authorization value"
// @Param body-payload body dto.CreateProjectRequest true "dto.CreateProjectRequest"
// @Success 200 {object} pkg.BaseResponse{data=dto.CreateProjectResponse} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/projects [post]
func (h *ProjectHandler) CreateProject() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		var err error
		var payloadReq dto.CreateProjectRequest

		// Retrieve values from context (locals)
		phoneNumber, _ := r.Context().Value(middleware.PhoneNumber).(string)

		decoder := json.NewDecoder(r.Body)
		if err = decoder.Decode(&payloadReq); err != nil {
			httperror.SetResponse(w, 400, "body payload required")
			return
		}
		// validate payload
		errMsgs := pkg.ValidateStruct(payloadReq)
		if len(errMsgs) > 0 {
			httperror.SetResponse(w, 400, errMsgs)
			return
		}
		defer r.Body.Close()

		project, err := h.projectUsecase.CreateNewInquiry(r.Context(), phoneNumber, projectDomain.Entity{
			Name:        payloadReq.Name,
			Description: payloadReq.Description,
			ServiceType: payloadReq.ServiceType,
			TargetTime:  payloadReq.TimePriority,
		})
		if err != nil {
			log.Println("error creating projects: ", string(debug.Stack()))
			httperror.SetResponse(w, 500, "internal server error")
			return
		}

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data: dto.CreateProjectResponse{
				UUID: project.UUID,
			},
		})

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(dataByte)
		if err != nil {
			return
		}
	}
}

// @Summary Update Project
// @Description Update Project
// @Tags Projects
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param authorization header string true "Authorization value"
// @Param body-payload body dto.UpdateProjectReq true "dto.UpdateProjectReq"
// @Param project_uuid path string true "uuid project"
// @Success 200 {object} pkg.BaseResponse{} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/projects/{project_uuid} [patch]
func (h *ProjectHandler) UpdateProject() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		uuid := ps.ByName("uuid")

		var payloadReq dto.UpdateProjectReq
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&payloadReq); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"code":400,"message":"body payload required"}`))
			return
		}
		// validate payload
		errMsgs := pkg.ValidateStruct(payloadReq)
		if len(errMsgs) > 0 {
			errByte, _ := json.Marshal(pkg.BaseResponse{
				Code:    400,
				Message: "error",
				Data:    errMsgs,
			})

			w.Header().Set("Content-Type", "application/json")
			w.Write(errByte)
			return
		}
		defer r.Body.Close()

		data := map[string]interface{}{
			"uuid":         uuid,
			"name":         payloadReq.Name,
			"description":  payloadReq.Description,
			"service_type": payloadReq.ServiceType,
			"status":       "Inquiry",
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

// @Summary Project Histories
// @Description Project Histories
// @Tags Projects
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param authorization header string false "Authorization value"
// @Param basic-param query commondto.Pagination true "basic param"
// @Param project_uuid path string true "project_uuid"
// @Success 200 {object} pkg.BaseResponse{data=[]dto.ProjectHistoryRes} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/projects/{project_uuid}/histories [get]
func (h *ProjectHandler) ProjectHistories() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		// Retrieve values from context (locals)
		phoneNumber, _ := r.Context().Value(middleware.PhoneNumber).(string)
		if phoneNumber == "" {
			log.Println("error getting project: ", string(debug.Stack()))
			httperror.SetResponse(w, 403, "forbidden")
			return
		}

		UUID := ps.ByName("uuid")

		var err error
		var payloadReq commondto.Pagination

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
			payloadReq.PerPage = 50
		}

		data, err := h.projectUsecase.ListProjectHistory(r.Context(), UUID, payloadReq.Page, payloadReq.PerPage)
		if err != nil {
			log.Println("error getting project history: ", string(debug.Stack()))
			httperror.SetResponse(w, 500, "internal server error")
			return
		}

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data:    data,
		})

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(dataByte)
		if err != nil {
			return
		}
	}
}
