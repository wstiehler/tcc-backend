package models

import (
	config "github.com/wstiehler/tcc-backend/internal/infrastructure/database"

	_ "github.com/go-sql-driver/mysql"
)

type vacancyRepository struct {
	Vacancies []VacancyEntity
}

func NewRepository() *vacancyRepository {
	return &vacancyRepository{}
}

func (c *vacancyRepository) SumGrossRevenue() (float64, error) {

	var statistics float64
	xpto := config.DB.Table("vacancies").Select("sum(gross_revenue")

	if xpto.Error != nil {
		return 0, xpto.Error
	}
	return statistics, nil
}

func (c *vacancyRepository) GetAll() ([]VacancyEntity, error) {
	config.DB.Find(&c.Vacancies)
	return c.Vacancies, nil
}

func (c *vacancyRepository) GetByID(id string) (VacancyEntity, error) {
	var vacancy VacancyEntity
	config.DB.Where("vacancy_id = ?", id).First(&vacancy)
	return vacancy, nil
}

func (c *vacancyRepository) Create(vacancy *VacancyEntity) (VacancyEntity, error) {
	config.DB.Create(&vacancy)
	return *vacancy, nil
}

func (c *vacancyRepository) Update(vacancy *VacancyEntity, vacancy_id string) (err error) {
	config.DB.Save(&vacancy)
	return nil
}

func (c *vacancyRepository) Delete(vacancy *VacancyEntity, vacancy_id string) error {
	config.DB.Where("vacancy_id = ?", vacancy_id).Delete(&VacancyEntity{})
	return nil
}

func (c *vacancyRepository) Count() int {
	var count int
	config.DB.Model(&VacancyEntity{}).Count(&count)
	return count
}
