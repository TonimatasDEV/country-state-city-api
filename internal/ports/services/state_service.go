package services

import (
	"country-state-city-api/internal/domain"
	"country-state-city-api/internal/ports/repositories"
)

type StateService struct {
	countryRepo repositories.CountryRepository
}

func NewStateService(countryRepo repositories.CountryRepository) *StateService {
	return &StateService{countryRepo: countryRepo}
}

func (service *StateService) GetStates(countryCode string) []domain.State {
	return service.countryRepo.GetStates(countryCode)
}
