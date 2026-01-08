package services

import (
	"country-state-city-api/internal/ports/repositories"
)

type StateService struct {
	countryRepo repositories.CountryRepository
}

func NewStateService(countryRepo repositories.CountryRepository) *StateService {
	return &StateService{countryRepo: countryRepo}
}

func (service *StateService) GetStatesNames(countryName string) []string {
	var states []string

	for _, state := range service.countryRepo.GetStates(countryName) {
		states = append(states, state.Name)
	}

	return states
}
