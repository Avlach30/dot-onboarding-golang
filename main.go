package main

import (
	"encoding/json"
	handler2 "github.com/codespace-id/codespace-x/app/handler"
	"github.com/codespace-id/codespace-x/app/repository"
	"github.com/codespace-id/codespace-x/app/usecase"
	"github.com/codespace-id/codespace-x/config"
	"github.com/codespace-id/codespace-x/pkg/dbconn/mysql"
	"log"
	"net/http"

	_ "github.com/codespace-id/codespace-x/docs"
	"github.com/codespace-id/codespace-x/pkg"
	"github.com/julienschmidt/httprouter"
	"github.com/swaggo/http-swagger"
)

// @title Codespace X REST API
// @version 1.0
// @description Codespace X
// @contact.name Codespace Indonesia
// @contact.url https://codespace.id
// @contact.email mail@codespace.id
func main() {
	router := httprouter.New()

	// instance config
	config.InitConfig()

	dbEnv := config.GetMySqlEnv()
	db, err := mysql.NewMysqlDB(dbEnv.Host, dbEnv.Username, dbEnv.Password, dbEnv.Db)
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
	handler2.NewUserHandler(router, userUsecase)
	handler2.NewAuthHandler(router, userUsecase)
	handler2.NewBannerHandler(router)
	handler2.NewNotificationHandler(router)
	handler2.NewProjectHandler(router)

	log.Println("=== SERVER STARTED at PORT 7777 ===")
	log.Fatal(http.ListenAndServe(":7777", router))
}
