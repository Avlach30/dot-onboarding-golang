package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codespace-id/codespace-x/pkg"
	"github.com/julienschmidt/httprouter"
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

func (h *ProjectHandler) ListProject() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		data := []map[string]interface{}{
			{
				"uuid":         "CYUS-898H",
				"name":         "Test TOEFL Online",
				"description":  "Dapatkan wordpress landing",
				"service_type": "web apps development",
				"status":       "On Going",
				"created_at":   "2011-08-12T20:17:46.384Z",
				"astrodevs": []map[string]interface{}{
					{
						"fullname":  "Hiegar",
						"role":      "UI/UX",
						"image_url": "https://res.cloudinary.com/deafomwc7/image/upload/v1664837475/codespace/images/team/team-4a_aqfwhw.jpg",
					},
					{
						"fullname":  "Ubai",
						"role":      "Backend",
						"image_url": "https://res.cloudinary.com/deafomwc7/image/upload/v1664837474/codespace/images/team/team-1a_asflru.jpg",
					},
				},
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

func (h *ProjectHandler) DetailProject() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		uuid := ps.ByName("uuid")
		data := map[string]interface{}{
			"uuid":         uuid,
			"name":         "Test TOEFL Online",
			"description":  "Dapatkan wordpress landing",
			"service_type": "web apps development",
			"status":       "On Going",
			"created_at":   "2011-08-12T20:17:46.384Z",
			"astrodevs": []map[string]interface{}{
				{
					"fullname":  "Hiegar",
					"role":      "UI/UX",
					"image_url": "https://res.cloudinary.com/deafomwc7/image/upload/v1664837475/codespace/images/team/team-4a_aqfwhw.jpg",
				},
				{
					"fullname":  "Ubai",
					"role":      "Backend",
					"image_url": "https://res.cloudinary.com/deafomwc7/image/upload/v1664837474/codespace/images/team/team-1a_asflru.jpg",
				},
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
