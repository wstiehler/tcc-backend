package models

//Create a repository in memory for test unitary

import (
	"sync"
)

type vacancyRepositoryMemory struct {
	Vacancies []VacancyEntity
	mu        sync.Mutex
}

func NewRepositoryMemory() *vacancyRepositoryMemory {
	return &vacancyRepositoryMemory{
		Vacancies: []VacancyEntity{},
	}
}

func (r *vacancyRepositoryMemory) CreateVacancy(vacancy *VacancyEntity) (err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Vacancies = append(r.Vacancies, *vacancy)
	return nil
}

func (r *vacancyRepositoryMemory) GetVacancyByID(id string) (VacancyEntity, error) {

	for _, item := range r.Vacancies {
		if item.Id == id {
			return item, nil
		}
	}
	return VacancyEntity{}, nil
}

func (r *vacancyRepositoryMemory) UpdateVacancy(vacancy *VacancyEntity, id string) (VacancyEntity, error) {

	for i, item := range r.Vacancies {
		if item.Id == id {
			r.Vacancies[i] = *vacancy
			return *vacancy, nil
		}
	}
	return VacancyEntity{}, nil
}

func (r *vacancyRepositoryMemory) DeleteVacancy(id string) error {

	for i, item := range r.Vacancies {
		if item.Id == id {
			r.Vacancies = append(r.Vacancies[:i], r.Vacancies[i+1:]...)
			return nil
		}
	}
	return nil
}

func (r *vacancyRepositoryMemory) Count() int {
	return len(r.Vacancies)
}
