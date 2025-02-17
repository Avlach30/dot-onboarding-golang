package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"os"
	"strings"
	"time"

	"fmt"
	"log"
	"strconv"

	handler "gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/handler"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/seeder"
	"gorm.io/gorm"

	userRepo "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/repository"
	userUC "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/usecase"

	roleRepo "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/repository"
	roleUC "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/usecase"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/migration"

	permissionRepo "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/repository"
	permissionUC "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/usecase"

	authJob "gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/job"
	authRepo "gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/repository"
	authUC "gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/usecase"
	fileUC "gitlab.dot.co.id/playground/boilerplates/golang-service/app/storage/usecase"

	notificationRepo "gitlab.dot.co.id/playground/boilerplates/golang-service/app/notification/repository"
	notificationUC "gitlab.dot.co.id/playground/boilerplates/golang-service/app/notification/usecase"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/dbconn"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/storage"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/task"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
)

var (
	// global variable
	db             *gorm.DB
	router         *gin.Engine
	workers        *task.Workers
	storageManager storage.StorageManager

	// args variable
	execMigration     *string
	runMigration      *string
	onlyJobExecutor   *string
	withJobExecutor   *string
	migrationFileName *string
	runSeeder         *string
	seederClass       *string
)

func main() {
	initializeLog()

	extractArgs()

	if err := initializeDatabase(); err != nil {
		panic(err)
	}

	handleMigrationAndSeeding()

	initializeSingleton()

	if *withJobExecutor == "true" || *onlyJobExecutor == "true" {
		initializeWorkers()
	}

	if err := initializeStorageManager(); err != nil {
		panic(err)
	}

	if *onlyJobExecutor != "true" {
		initializeRouter()
		initializeModule()
		initializeSentry()
		startHttpServer()
	} else {
		startIdleServer()
	}
}

func startIdleServer() {
	if *onlyJobExecutor == "true" {
		log.Println("Only job executor mode")
		for {
			time.Sleep(time.Second)
		}
	}
}

func startHttpServer() {
	log.Printf("Starting server on port %s...", config.AppPort)
	if err := router.Run(":" + config.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func extractArgs() {
	// migration args requirement
	execMigration = flag.String("exec", "up", "Args (e.g., --migration [up/down/fresh]")
	runMigration = flag.String("migration", "false", "Args (e.g., --migration [true/false]")
	migrationFileName = flag.String("fileName", "", "Migration file name")

	// seeder args requirement
	runSeeder = flag.String("dbseed", "false", "")
	seederClass = flag.String("class", "", "Add multiple values (e.g., --class UserSeed,RoleSeeder,...)")

	// job executor args requirement
	onlyJobExecutor = flag.String("onlyJobExecutor", "false", "true for only job executor, otherwise false")
	withJobExecutor = flag.String("withJobExecutor", "false", "true for run server with job executor, otherwise false")

	flag.Parse()
}

func initializeDatabase() error {
	var err error
	db, err = dbconn.InitDb(&dbconn.DatabaseCredentials{
		Host:     config.DBHost,
		Username: config.DBUsername,
		Password: config.DBPassword,
		Port:     config.DBPort,
		Name:     config.DBName,
		TimeZome: config.DBTimeZone,
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func initializeLog() {
	logWriter := utils.NewLogWriter()
	log.SetOutput(logWriter)
}

func handleMigrationAndSeeding() {
	if *runMigration == "true" {
		if *execMigration == "create" {
			migration.Create(db, *migrationFileName)
		} else {
			migration.Run(db, *execMigration)
		}
		os.Exit(0)
	}

	if *runSeeder == "true" {
		var classes []string = nil
		if *seederClass != "" {
			classes = strings.Split(*seederClass, ",")
		}

		if err := seeder.Run(db, classes); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
}

func initializeStorageManager() error {
	var err error

	log.Printf("Using storage type: %s\n", config.Storage)
	switch config.Storage {
	case "gcs":
		storageManager, err = storage.NewGCSManager()
	case "s3":
		storageManager, err = storage.NewS3Manager()
	case "minio":
		storageManager, err = storage.NewMinIOManager()
	default:
		log.Printf("Invalid storage type: %s\n", config.Storage)
	}

	if err != nil {
		log.Printf("Failed to initialize storage manager: %v\n", err)
	}

	return err
}

func initializeSingleton() {
	singleton.InitGlobal(workers, db, &storageManager)
}

func initializeWorkers() {
	workers = task.InitQueueWorkerTask()

	if *withJobExecutor == "true" || *onlyJobExecutor == "true" {

		singleton.AddJobDictionary(authJob.Jobs())

		go singleton.ExecuteJobTask()

		schedulerExecutor := task.InitAllSchedulerTask()
		go schedulerExecutor.RunScheduler()
	}

	go task.RunAllActiveWorker(workers)
}

func initializeRouter() {
	router = gin.New()

	gin.SetMode(config.GinMode)

	router.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	router.Use(handler.RecoverPanic())

	healthCheck(router)
}

func initializeSentry() {
	tracesSampleRate, _ := strconv.ParseFloat(config.SentrySampleTrace, 64)
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.SentryDSN,
		EnableTracing:    true,
		TracesSampleRate: tracesSampleRate,
		Debug:            true,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	} else {
		fmt.Println("Sentry initialized")
	}
}

func initializeModule() {
	// Initialize repositories
	userRepo := userRepo.NewUserRepository(db)
	roleRepo := roleRepo.NewRoleRepository(db)
	permissionRepo := permissionRepo.NewPermissionRepository(db)
	authRepo := authRepo.NewAuthRepository(db)
	notificationRepo := notificationRepo.NewNotificationRepository(db)

	// Initialize usecases
	userUsecase := userUC.NewUserUsecase(userRepo)
	permissionUsecase := permissionUC.NewPermissionUsecase(permissionRepo)
	roleUsecase := roleUC.NewRoleUsecase(roleRepo)
	authUsecase := authUC.NewAuthUsecase(authRepo)
	fileUsecase := fileUC.NewFileUsecase()
	notificationUsecase := notificationUC.NewNotificationUseCase(notificationRepo)

	// Setup handlers
	handler.NewUserHandler(router, userUsecase)
	handler.NewPermissionHandler(router, permissionUsecase)
	handler.NewRoleHandler(router, roleUsecase)
	handler.NewAuthHandler(router, authUsecase)
	handler.NewCommonHandler(router, fileUsecase)
	handler.NewNotificationHandler(router, notificationUsecase)
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
