package models

import (
	config "github.com/wstiehler/tcc-backend/internal/infrastructure/database"

	_ "github.com/go-sql-driver/mysql"
)

type vacancyRepository struct {
	Vacancies []VacancyEntity
}

type responsibilyRepository struct {
	Responsible []ResponsibleEntity
}

func NewRepository() *vacancyRepository {
	return &vacancyRepository{}
}

func NewRepositoryResponsibily() *responsibilyRepository {
	return &responsibilyRepository{}
}

func (c *vacancyRepository) GetAll() ([]VacancyEntity, error) {
	config.DB.Find(&c.Vacancies)
	return c.Vacancies, nil
}

func (c *vacancyRepository) GetByID(id string) (VacancyEntity, error) {
	var vacancy VacancyEntity
	config.DB.Where("id = ?", id).First(&vacancy)
	return vacancy, nil
}

func (r *responsibilyRepository) GetByID(id string) (ResponsibleEntity, error) {
	var responsible ResponsibleEntity
	config.DB.Where("id = ?", id).First(&responsible)
	return responsible, nil
}

func (c *vacancyRepository) Create(vacancy *VacancyEntity) (VacancyEntity, error) {
	config.DB.Create(&vacancy)
	return *vacancy, nil
}

func (r *responsibilyRepository) CreateResponsibily(responsible *ResponsibleEntity) (ResponsibleEntity, error) {
	config.DB.Create(&responsible)
	return *responsible, nil
}

func (c *vacancyRepository) Update(vacancy *VacancyEntity, id string) (err error) {
	config.DB.Save(&vacancy)
	return nil
}

func (c *vacancyRepository) Delete(vacancy *VacancyEntity, id string) error {
	config.DB.Where("id = ?", id).Delete(&VacancyEntity{})
	return nil
}

func (c *vacancyRepository) Count() int {
	var count int
	config.DB.Model(&VacancyEntity{}).Count(&count)
	return count
}
