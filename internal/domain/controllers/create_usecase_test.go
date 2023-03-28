package controllers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	models "github.com/wstiehler/tcc-backend/internal/domain/models"
)

func TestVacancy(t *testing.T) {

	t.Run("Create a vacancy", func(t *testing.T) {
		vacancy := models.VacancyEntity{
			VacancyName: "Evoluinfo",
		}

		err := models.NewRepositoryMemory().CreateVacancy(&vacancy)

		assert.Nil(t, err)
	})

	t.Run("Get a vacancy by id", func(t *testing.T) {
		vacancy, err := models.NewRepositoryMemory().GetVacancyByID("1")

		assert.Nil(t, err)
		assert.NotNil(t, vacancy)
	})

	t.Run("Update a vacancy", func(t *testing.T) {
		vacancy := models.VacancyEntity{
			VacancyName: "Evoluinfo",
		}

		vacancy, err := models.NewRepositoryMemory().UpdateVacancy(&vacancy, "1")

		assert.Nil(t, err)

		assert.NotNil(t, vacancy)
	})

	t.Run("Delete a vacancy", func(t *testing.T) {
		err := models.NewRepositoryMemory().DeleteVacancy("1")

		assert.Nil(t, err)
	})

}
