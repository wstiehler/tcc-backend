package models

import (
	"time"
)

type ModelsEntity struct {
	VacancyEntity     *VacancyEntity
	ResponsibleEntity *ResponsibleEntity
}

type VacancyEntity struct {
	Id          string `json:"id" gorm:"primary_key"`
	VacancyName string `json:"vacancy_name"`

	Salary          float64 `json:"salary"`
	IsHomeOffice    string  `json:"is_home_office"`
	LevelExperience string  `json:"level_experience"`

	CulturalCaracteristics   string `json:"cultural_caracteristics"`
	HasPhysicalAccessibility string `json:"has_physical_accessibility"`
	HasSuperiorMonitors      string `json:"has_superior_monitors"`
	Specification            string `json:"specification" `
	Description              string `json:"description"`

	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	CompanyName string `json:"company_name"`
	Email       string `json:"email"`

	ResponsibleId string

	Responsible ResponsibleEntity `json:"responsible" gorm:"foreignkey:ResponsibleId"`
}

type ResponsibleEntity struct {
	Id              string `json:"id" gorm:"primary_key"`
	ResponsibleName string `json:"responsible_name"`
	Contact         string `json:"contact"`
}

func (b *VacancyEntity) TableName() string {
	return "vacancies"
}

func (b *ResponsibleEntity) TableName() string {
	return "responsibles"
}
