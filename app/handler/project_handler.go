package handler

import (
	"encoding/json"
	"github.com/codespace-id/codespace-x/app/dto"
	"github.com/codespace-id/codespace-x/pkg"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ProjectHandler struct {
}

func NewProjectHandler(router *httprouter.Router) {
	basePath := "/api/v1/projects"
	projectHandler := &ProjectHandler{}

	router.GET(basePath, projectHandler.ListProject())
	router.GET(basePath+"/:uuid", projectHandler.DetailProject())
	router.POST(basePath, projectHandler.CreateProject())
	router.PATCH(basePath+"/:uuid", projectHandler.UpdateProject())

}

// @Summary List Project
// @Description List Project
// @Tags Projects
// @Accept json
// @Produce json
// @Param authorization header string false "Authorization value"
// @Param basic-param query pkg.Pagination true "basic param"
// @Success 200 {object} pkg.BaseResponse{data=[]dto.ListProjectResponse} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/projects [get]
func (h *ProjectHandler) ListProject() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data: []dto.ListProjectResponse{
				{
					UUID:        "CYUS-898H",
					Name:        "Test TOEFL Online",
					Description: "Dapatkan wordpress landing",
					ServiceType: "web apps development",
					Status:      "On Going",
					CreatedAt:   "2011-08-12T20:17:46.384Z",
					Astrodevs: []dto.UserResponse{
						{
							Fullname: "Hiegar",
							Role:     "UI/UX",
							ImageURL: "https://res.cloudinary.com/deafomwc7/image/upload/v1664837475/codespace/images/team/team-4a_aqfwhw.jpg",
						},
						{
							Fullname: "Ubai",
							Role:     "Backend",
							ImageURL: "https://res.cloudinary.com/deafomwc7/image/upload/v1664837474/codespace/images/team/team-1a_asflru.jpg",
						},
					},
				},
			},
			Meta: &pkg.MetaResponse{
				Page:  1,
				Limit: 10,
			},
		})

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(dataByte)
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
// @Param authorization header string false "Authorization value"
// @Param project_uuid path string true "project_uuid"
// @Success 200 {object} pkg.BaseResponse{data=dto.ListProjectResponse} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/projects/{project_uuid} [get]
func (h *ProjectHandler) DetailProject() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		uuid := ps.ByName("uuid")

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data: dto.ListProjectResponse{

				UUID:        uuid,
				Name:        "Test TOEFL Online",
				Description: "Dapatkan wordpress landing",
				ServiceType: "web apps development",
				Status:      "On Going",
				CreatedAt:   "2011-08-12T20:17:46.384Z",
				Astrodevs: []dto.UserResponse{
					{
						Fullname: "Hiegar",
						Role:     "UI/UX",
						ImageURL: "https://res.cloudinary.com/deafomwc7/image/upload/v1664837475/codespace/images/team/team-4a_aqfwhw.jpg",
					},
					{
						Fullname: "Ubai",
						Role:     "Backend",
						ImageURL: "https://res.cloudinary.com/deafomwc7/image/upload/v1664837474/codespace/images/team/team-1a_asflru.jpg",
					},
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

func (h *ProjectHandler) CreateProject() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data: map[string]interface{}{
				"uuid": "XTY9A-YUABQ",
			},
		})

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(dataByte)
		if err != nil {
			return
		}
	}
}

func (h *ProjectHandler) UpdateProject() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		uuid := ps.ByName("uuid")

		// payload
		type payload struct {
			ServiceType  int    `json:"service_type"`
			Name         string `json:"name"`
			Description  string `json:"description"`
			DeadlineType int    `json:"deadline_type"`
		}
		var payloadReq payload
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
