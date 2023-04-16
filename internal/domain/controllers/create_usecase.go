package controllers

import (
	"net/http"

	models "github.com/wstiehler/tcc-backend/internal/domain/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	var responsible models.ResponsibleEntity

	vacancies, err := models.NewRepository().GetAll()

	for i := 0; i < len(vacancies); i++ {
		responsible, _ = models.NewRepositoryResponsibily().GetByID(responsible.Id)
		vacancies[i].Responsible.ResponsibleName = responsible.ResponsibleName
		vacancies[i].Responsible.Contact = responsible.Contact
		vacancies[i].Responsible.Id = responsible.Id

	}

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"vacancies": vacancies})
	}
}

func GetVacanciesActive(c *gin.Context) {
	var vacancies []models.VacancyEntity

	var responsible models.ResponsibleEntity

	vacancies, err := models.NewRepository().GetAllActive()

	for i := 0; i < len(vacancies); i++ {
		responsible, _ = models.NewRepositoryResponsibily().GetByID(responsible.Id)
		vacancies[i].Responsible.ResponsibleName = responsible.ResponsibleName
		vacancies[i].Responsible.Contact = responsible.Contact
		vacancies[i].Responsible.Id = responsible.Id
	}

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"vacancies": vacancies})
	}
}

func GetVacanciesByEmail(c *gin.Context) {
	email := c.Param("email")

	var vacancies []models.VacancyEntity

	var responsible models.ResponsibleEntity

	vacancies, err := models.NewRepository().GetByEmail(email)

	for i := 0; i < len(vacancies); i++ {
		responsible, _ = models.NewRepositoryResponsibily().GetByID(responsible.Id)
		vacancies[i].Responsible.ResponsibleName = responsible.ResponsibleName
		vacancies[i].Responsible.Contact = responsible.Contact
		vacancies[i].Responsible.Id = responsible.Id
	}

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"vacancies": vacancies})
	}
}

func CreateVacancy(c *gin.Context) {
	var vacancy models.VacancyEntity

	vacancy.Id = uuid.New().String()

	c.BindJSON(&vacancy)

	vacancy.Active = true
	vacancy.Responsible.Id = vacancy.Id

	vacancy, err := models.NewRepository().Create(&vacancy)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, vacancy)
	}
}

func GetVacancyByID(c *gin.Context) {
	id := c.Param("id")

	var vacancy models.VacancyEntity

	var responsible models.ResponsibleEntity

	vacancy, err := models.NewRepository().GetByID(id)

	responsible, _ = models.NewRepositoryResponsibily().GetByID(id)

	//TODO: Create a struct to return the
	vacancy.Responsible.Id = responsible.Id
	vacancy.Responsible.ResponsibleName = responsible.ResponsibleName
	vacancy.Responsible.Contact = responsible.Contact

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)

	} else {
		c.JSON(http.StatusOK, vacancy)
	}
}

func GetResponsibleByID(c *gin.Context) {
	id := c.Param("id")

	var responsible models.ResponsibleEntity

	responsible, err := models.NewRepositoryResponsibily().GetByID(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, responsible)
	}
}

func UpdateVacancy(c *gin.Context) {
	var vacancy models.VacancyEntity

	id := c.Param("id")

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

func InactivateVacancy(c *gin.Context) {
	var vacancy models.VacancyEntity

	id := c.Param("id")

	vacancy, err := models.NewRepository().GetByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, vacancy)
	}

	vacancy.Active = false

	err = models.NewRepository().Update(&vacancy, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"inactivate": vacancy})
	}

}

func DeleteVacancy(c *gin.Context) {
	var vacancy models.VacancyEntity

	id := c.Param("id")

	err := models.NewRepository().Delete(&vacancy, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"deleted": id})
	}

}
