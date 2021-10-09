package app

import (
	"log"
	"os"
	"regexp"

	"github.com/bookstores-go-microservices/users-api/db"
	"github.com/bookstores-go-microservices/users-api/routes"
	"github.com/bookstores-go-microservices/users-api/untils/loggers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

const projectDirName = "users-api" // change to relevant project name

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func StartApplication() {
	loadEnv()

	loggers.NewLogger()
	loggers.Info("Starting application...")

	// Db configuration
	db.NewDB(os.Getenv(db.DB_USER), os.Getenv(db.DB_PASSWD), os.Getenv(db.DB_HOST), os.Getenv(db.DB_SCHEME))

	v1 := router.Group("/api/v1")
	{
		routes.SetupRouter(v1)
	}
	router.Run(":8080")
}
