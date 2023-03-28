package models

import (
	"time"
)

type ModelsEntity struct {
	VacancyEntity     *VacancyEntity
	ResponsibleEntity *ResponsibleEntity
}

type VacancyEntity struct {
	VacancyId   uint64 `json:"vacancy_id" gorm:"primary_key"`
	VacancyName string `json:"vacancy_name"`

	Salary          float64 `json:"salary"`
	IsHomeOffice    bool    `json:"is_home_office"`
	LevelExperience string  `json:"level_experience"`

	CulturalCaracteristics   string `json:"cultural_caracteristics"`
	HasPhysicalAccessibility bool   `json:"has_physical_accessibility"`
	HasSuperiorMonitors      bool   `json:"has_superior_monitors"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	ResponsibleEntity *ResponsibleEntity `json:"responsible_entity"`
}

type ResponsibleEntity struct {
	CompanyName     string `json:"company_name"`
	ResponsibleName string `json:"responsible_name"`
	Email           string `json:"email"`
	Contact         string `json:"contact"`
}

type StatisticsvacanciesEntity struct {
	Count        int     `json:"count"`
	GrossRevenue float64 `json:"gross_revenue"`
}

func (b *VacancyEntity) TableName() string {
	return "vacancies"
}

func (b *ResponsibleEntity) TableName() string {
	return "responsibles"
}
