package main

import (
	"encoding/json"
	authHandler "github.com/codespace-id/codespace-x/app/auth/handler"
	"github.com/codespace-id/codespace-x/app/auth/repository"
	authUC "github.com/codespace-id/codespace-x/app/auth/usecase"
	"github.com/codespace-id/codespace-x/app/banner/handler"
	bannerRepo "github.com/codespace-id/codespace-x/app/banner/repository"
	bannerUC "github.com/codespace-id/codespace-x/app/banner/usecase"
	commonHandler "github.com/codespace-id/codespace-x/app/common/handler"
	"github.com/codespace-id/codespace-x/app/common/repository"
	notifHandler "github.com/codespace-id/codespace-x/app/notification/handler"
	projectHandler "github.com/codespace-id/codespace-x/app/project/handler"
	projectRepo "github.com/codespace-id/codespace-x/app/project/repository"
	projectUC "github.com/codespace-id/codespace-x/app/project/usecase"
	tncHandler "github.com/codespace-id/codespace-x/app/tnc/handler"
	userHandler "github.com/codespace-id/codespace-x/app/user/handler"
	userRepo "github.com/codespace-id/codespace-x/app/user/repository"
	userUC "github.com/codespace-id/codespace-x/app/user/usecase"
	"github.com/codespace-id/codespace-x/pkg/common/enum"
	"github.com/codespace-id/codespace-x/pkg/dbconn"
	"log"
	"net/http"

	"github.com/codespace-id/codespace-x/config"
	_ "github.com/codespace-id/codespace-x/docs"
	"github.com/codespace-id/codespace-x/pkg"
	"github.com/codespace-id/codespace-x/pkg/Integrations/otp/implementations/zenziva"
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

	db, err := dbconn.GetDb(enum.MYSQL)
	if err != nil {
		log.Fatal(err)
	}

	// 3rd parties
	zenzivaOTP := zenziva.NewZenziva(config.ZenzivaBaseURL, config.ZenzivaPassKey, config.ZenzivaUserKey)

	// repository
	userRepository := userRepo.NewUserRepository(db)
	otpRepo := repository.NewOtpRepository(db)
	bannerRepository := bannerRepo.NewBannerRepository(db)
	projectRepository := projectRepo.NewProjectRepository(db)
	sqlTxRepo := commonrepo.NewSqlTx(db)
	userProjectRepo := projectRepo.NewUserProjectRepository(db)
	projectImagesRepo := projectRepo.NewProjectImagesRepository(db)
	projectHistoryRepo := projectRepo.NewProjectHistoryRepository(db)

	// usecase
	userUsecase := userUC.NewUserUsecase(userRepository)
	authUsecase := authUC.NewAuthUsecase(zenzivaOTP, otpRepo, userRepository)
	bannerUsecase := bannerUC.NewBannerUsecase(bannerRepository)
	projectUsecase := projectUC.NewProjectUsecase(projectRepository, sqlTxRepo, userProjectRepo, userRepository, projectImagesRepo, projectHistoryRepo)
	projectPublicUsecase := projectUC.NewProjectPublicUsecase(projectRepository, sqlTxRepo, userProjectRepo, userRepository)

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

	userHandler.NewUserHandler(router, userUsecase)
	authHandler.NewAuthHandler(router, userUsecase, authUsecase)
	handler.NewBannerHandler(router, bannerUsecase)
	notifHandler.NewNotificationHandler(router)
	projectHandler.NewProjectHandler(router, projectUsecase, projectPublicUsecase)
	commonHandler.NewCommonHandler(router)
	tncHandler.NewTncHandler(router)

	log.Println("=== SERVER STARTED at PORT 7777 ===")
	log.Fatal(http.ListenAndServe(":7777", router))
}
