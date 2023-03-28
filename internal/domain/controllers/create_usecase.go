package controllers

import (
	"net/http"

	models "github.com/wstiehler/tcc-backend/internal/domain/models"

	"github.com/gin-gonic/gin"
)

type Creator interface {
	Createvacancy(c *gin.Context)
}

func EntityVacancy() *models.VacancyEntity {
	var vacancy models.VacancyEntity
	return &vacancy
}

func GetVacancies(c *gin.Context) {
	var vacancies []models.VacancyEntity

	vacancies, err := models.NewRepository().GetAll()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"vacancies": vacancies})
	}
}

func CreateVacancy(c *gin.Context) {
	var vacancy models.VacancyEntity

	c.BindJSON(&vacancy)

	vacancy, err := models.NewRepository().Create(&vacancy)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, vacancy)
	}
}

func GetVacancyByID(c *gin.Context) {
	id := c.Param("vacancy_id")

	var vacancy models.VacancyEntity

	vacancy, err := models.NewRepository().GetByID(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, vacancy)
	}
}

func UpdateVacancy(c *gin.Context) {
	var vacancy models.VacancyEntity

	id := c.Param("vacancy_id")

	vacancy, err := models.NewRepository().GetByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, vacancy)
	}

	c.BindJSON(&vacancy)

	if err != nil {
		c.JSON(http.StatusNotFound, "error calculating billing value")
	}

	err = models.NewRepository().Update(&vacancy, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"updated": vacancy})
	}

}

func DeleteVacancy(c *gin.Context) {
	var vacancy models.VacancyEntity

	id := c.Param("vacancy_id")

	err := models.NewRepository().Delete(&vacancy, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"deleted": id})
	}

}

func StatisticsVacancies(c *gin.Context) {
	var statistics models.StatisticsvacanciesEntity

	statistics.Count = models.NewRepository().Count()

	// statistics.GrossRevenue, _ = models.NewRepository().SumGrossRevenue()

	c.JSON(http.StatusOK, gin.H{"statistics": statistics})
}
