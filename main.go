package main

import (
	"encoding/json"
	"flag"

	"fmt"
	"log"
	"net/http"
	"strconv"

	handler "gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/handler"

	userRepo "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/repository"
	userUC "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/usecase"

	roleRepo "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/repository"
	roleUC "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/usecase"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/migration"

	permissionRepo "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/repository"
	permissionUC "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/usecase"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/gin-gonic/gin"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/dbconn"
)

func main() {
	db, err := dbconn.InitDb(
		&dbconn.DatabaseCredentials{
			Host:     config.DBHost,
			Username: config.DBUsername,
			Password: config.DBPassword,
			Port:     config.DBPort,
			Name:     config.DBName,
			TimeZome: config.DBTimeZone,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// migration run
	runMigration := flag.String("migration", "none", "--")
	autoMigration := flag.Bool("auto", false, "--")
	flag.Parse()
	if *runMigration == "true" {
		migration.Run(db, *autoMigration)
		return
	}

	router := gin.New()

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

	healthCheck(router)

	// repository
	userRepository := userRepo.NewUserRepository(db)
	roleRepository := roleRepo.NewRoleRepository(db)
	permissionRepository := permissionRepo.NewPermissionRepository(db)

	// usecase
	userUsecase := userUC.NewUserUsecase(userRepository)
	permissionUsecase := permissionUC.NewPermissionUsecase(permissionRepository)
	roleUsecase := roleUC.NewRoleUsecase(roleRepository)

	// middware at main.go
	http.Handle("/", sentryHandler.Handle(router))

	handler.NewUserHandler(router, userUsecase)
	handler.NewPermissionHandler(router, permissionUsecase)
	handler.NewRoleHandler(router, roleUsecase)

	log.Println("=== SERVER STARTED at PORT 7777 ===")
	log.Fatal(http.ListenAndServe(":7777", router))
}

func healthCheck(router *gin.Engine) {
	router.GET("/health", func(httpContext *gin.Context) {
		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data: struct {
				Service string `json:"service"`
				Status  string `json:"status"`
			}{
				Service: "Codespace X",
				Status:  "Healthy",
			},
		})

		httpContext.Data(200, "Content-Type: application/json", dataByte)
	})
}
