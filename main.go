package main

import (
	"encoding/json"
	"flag"
	"net/http"

	"fmt"
	"log"
	"strconv"

	handler "gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/handler"

	userRepo "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/repository"
	userUC "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/usecase"

	roleRepo "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/repository"
	roleUC "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/usecase"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/migration"

	permissionRepo "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/repository"
	permissionUC "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/usecase"

	authRepo "gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/repository"
	authUC "gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/usecase"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
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
	execMigration := flag.String("exec", "up", "--")
	flag.Parse()
	if *runMigration == "true" {
		migration.Run(db, *execMigration)
		return
	}

	router := gin.New()
	gin.SetMode(config.GinMode)

	tracesSampleRate, _ := strconv.ParseFloat(config.SentrySampleTrace, 64)

	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.SentryDSN,
		EnableTracing:    true,
		TracesSampleRate: tracesSampleRate,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	} else {
		fmt.Println("Sentry initialized")
	}

	// Create an instance of sentryhttp
	sentryHandlerGin := sentrygin.New(sentrygin.Options{})

	// repository
	userRepository := userRepo.NewUserRepository(db)
	roleRepository := roleRepo.NewRoleRepository(db)
	permissionRepository := permissionRepo.NewPermissionRepository(db)
	authRepository := authRepo.NewAuthRepository(db)

	// usecase
	userUsecase := userUC.NewUserUsecase(userRepository)
	permissionUsecase := permissionUC.NewPermissionUsecase(permissionRepository)
	roleUsecase := roleUC.NewRoleUsecase(roleRepository)
	authUsecase := authUC.NewAuthUsecase(authRepository)

	// middware at main.go
	router.Use(sentryHandlerGin)
	router.Use(handler.Recovery500())
	healthCheck(router)

	handler.NewUserHandler(router, userUsecase)
	handler.NewPermissionHandler(router, permissionUsecase)
	handler.NewRoleHandler(router, roleUsecase)
	handler.NewAuthHandler(router, authUsecase)

	if err := router.Run(":" + config.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func healthCheck(router *gin.Engine) {
	router.GET("/health", func(httpContext *gin.Context) {
		dataByte, _ := json.Marshal(pkg.BaseResponse{
			StatusCode: http.StatusOK,
			Data:       nil,
		})

		httpContext.Data(http.StatusOK, "Content-Type: application/json", dataByte)
	})
}
