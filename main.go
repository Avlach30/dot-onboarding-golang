package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codespace-id/codespace-x/app/handler"
	"github.com/codespace-id/codespace-x/app/repository"
	"github.com/codespace-id/codespace-x/app/usecase"
	"github.com/codespace-id/codespace-x/config"
	"github.com/codespace-id/codespace-x/pkg/dbconn/mysql"

	_ "github.com/codespace-id/codespace-x/docs"
	"github.com/codespace-id/codespace-x/pkg"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Codespace X REST API
// @version 1.0
// @description Codespace X
// @contact.name Codespace Indonesia
// @contact.url https://codespace.id
// @contact.email mail@codespace.id
func main() {

	router := httprouter.New()

	db, err := mysql.NewMysqlDB(config.Host, config.Username, config.Password, config.Database)
	if err != nil {
		log.Fatal(err)
	}

	// repository
	userRepo := repository.NewUserRepository(db)

	// usecase
	userUsecase := usecase.NewUserUsecase(userRepo)

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		type healthRes struct {
			Service string `json:"service"`
			Status  string `json:"status"`
		}
		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data: healthRes{
				Service: "Codespace X",
				Status:  "Healthy",
			},
		})

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(dataByte)
		if err != nil {
			return
		}
	})

	router.GET("/swagger/*filepath", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		httpSwagger.WrapHandler(w, r)
	})
	handler.NewUserHandler(router, userUsecase)
	handler.NewAuthHandler(router, userUsecase)
	handler.NewBannerHandler(router)
	handler.NewNotificationHandler(router)
	handler.NewProjectHandler(router)

	log.Println("=== SERVER STARTED at PORT 7777 ===")
	log.Fatal(http.ListenAndServe(":7777", router))
}
