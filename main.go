package main

import (
	"encoding/json"
	authHandler "gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/handler"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/repository"
	authUC "gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/usecase"
	commonHandler "gitlab.dot.co.id/playground/boilerplates/golang-service/app/common/handler"

	// TODO: using this when need to use transaction
	// "gitlab.dot.co.id/playground/boilerplates/golang-service/app/common/repository"

	notifHandler "gitlab.dot.co.id/playground/boilerplates/golang-service/app/notification/handler"
	userHandler "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/handler"
	userRepo "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/repository"
	userUC "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/usecase"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/common/enum"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/dbconn"
	"log"
	"fmt"
	"strconv"
	"net/http"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	_ "gitlab.dot.co.id/playground/boilerplates/golang-service/docs"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/Integrations/otp/implementations/zenziva"
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
	
	tracesSampleRate, _ := strconv.ParseFloat(config.SentrySampleTrace, 64)

	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: config.SentryDSN,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production.
		TracesSampleRate: tracesSampleRate,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	
	// Create an instance of sentryhttp
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	// 3rd parties
	zenzivaOTP := zenziva.NewZenziva(config.ZenzivaBaseURL, config.ZenzivaPassKey, config.ZenzivaUserKey)

	// repository
	userRepository := userRepo.NewUserRepository(db)
	otpRepo := repository.NewOtpRepository(db)
	
	// prepare sql for transactions
	// sqlTxRepo := commonrepo.NewSqlTx(db)

	// usecase
	userUsecase := userUC.NewUserUsecase(userRepository)
	authUsecase := authUC.NewAuthUsecase(zenzivaOTP, otpRepo, userRepository)

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

	http.Handle("/", sentryHandler.Handle(router))
	
	userHandler.NewUserHandler(router, userUsecase)
	authHandler.NewAuthHandler(router, userUsecase, authUsecase)
	notifHandler.NewNotificationHandler(router)
	commonHandler.NewCommonHandler(router)


	log.Println("=== SERVER STARTED at PORT 7777 ===")
	log.Fatal(http.ListenAndServe(":7777", router))
}
