package web

import (
	"github.com/wstiehler/tcc-backend/internal/api/middlewares"
	"github.com/wstiehler/tcc-backend/internal/domain/controllers"
	"github.com/wstiehler/tcc-backend/internal/environment"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	env := environment.GetInstance()

	r := gin.Default()

	r.SetTrustedProxies([]string{env.APPLICATION_ADDRESS})

	r.Use(middlewares.CORSMiddleware())

	grp1 := r.Group("/v1")
	{
		grp1.GET("vacancies", controllers.GetVacancies)
		grp1.POST("vacancy", controllers.CreateVacancy)
		grp1.GET("vacancy/:id", controllers.GetVacancyByID)
		grp1.GET("vacancies/actives", controllers.GetVacanciesActive)
		grp1.GET("vacancies/:email", controllers.GetVacanciesByEmail)
		grp1.GET("responsible/:id", controllers.GetResponsibleByID)
		grp1.PUT("vacancy/:id", controllers.UpdateVacancy)
		grp1.PUT("vacancy/inactivate/:id", controllers.InactivateVacancy)
		grp1.DELETE("vacancy/:id", controllers.DeleteVacancy)
	}
	return r
}
