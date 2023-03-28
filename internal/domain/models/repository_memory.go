package models

//Create a repository in memory for test unitary

import (
	"strconv"
	"sync"
)

type vacancyRepositoryMemory struct {
	Companies []VacancyEntity
	mu        sync.Mutex
}

func NewRepositoryMemory() *vacancyRepositoryMemory {
	return &vacancyRepositoryMemory{
		Companies: []VacancyEntity{},
	}
}

func (r *vacancyRepositoryMemory) CreateVacancy(vacancy *VacancyEntity) (err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Companies = append(r.Companies, *vacancy)
	return nil
}

func (r *vacancyRepositoryMemory) GetVacancyByID(id string) (VacancyEntity, error) {

	vacancyID, _ := strconv.ParseUint(id, 10, 64)

	for _, item := range r.Companies {
		if item.VacancyId == vacancyID {
			return item, nil
		}
	}
	return VacancyEntity{}, nil
}

func (r *vacancyRepositoryMemory) UpdateVacancy(vacancy *VacancyEntity, id string) (VacancyEntity, error) {

	vacancyID, _ := strconv.ParseUint(id, 10, 64)

	for i, item := range r.Companies {
		if item.VacancyId == vacancyID {
			r.Companies[i] = *vacancy
			return *vacancy, nil
		}
	}
	return VacancyEntity{}, nil
}

func (r *vacancyRepositoryMemory) DeleteVacancy(id string) error {

	vacancyID, _ := strconv.ParseUint(id, 10, 64)

	for i, item := range r.Companies {
		if item.VacancyId == vacancyID {
			r.Companies = append(r.Companies[:i], r.Companies[i+1:]...)
			return nil
		}
	}
	return nil
}

func (r *vacancyRepositoryMemory) Count() int {
	return len(r.Companies)
}
