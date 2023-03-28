package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	web "github.com/wstiehler/tcc-backend/internal/api/handlers"
	"github.com/wstiehler/tcc-backend/internal/api/middlewares"
	"github.com/wstiehler/tcc-backend/internal/domain/models"
	"github.com/wstiehler/tcc-backend/internal/environment"
	config "github.com/wstiehler/tcc-backend/internal/infrastructure/database"
	"go.uber.org/zap"
)

var err error

var ListenAddr = ":8080"

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	env := environment.GetInstance()

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()

	config.DB.AutoMigrate(models.VacancyEntity{}, models.ResponsibleEntity{})

	if env.APPLICATION_PORT != "" {
		ListenAddr = env.APPLICATION_PORT
	}

	router := initRouter()

	err = router.Run(ListenAddr)

	handleErr(err)

}

func initRouter() *gin.Engine {
	r := web.SetupRouter()

	logger, err := zap.NewProduction()

	defer logger.Sync()

	handleErr(err)

	r.Use(middlewares.Logger(logger))
	r.Use(middlewares.Recovery(logger, true))
	r.Use(middlewares.CORSMiddleware())

	r.Run()

	return r
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
