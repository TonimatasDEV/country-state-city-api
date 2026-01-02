package services

import (
	"country-state-city-api/internal/ports/repositories"
	"strings"
)

type StateService struct {
	countryRepo repositories.CountryRepository
}

func NewStateService(countryRepo repositories.CountryRepository) *StateService {
	return &StateService{countryRepo: countryRepo}
}

func (service *StateService) GetStatesNames(countryName string) []string {
	countries := service.countryRepo.Get()

	var states []string

	for _, country := range countries {
		if strings.EqualFold(countryName, country.Name) {
			for _, state := range country.States {
				states = append(states, state.Name)
			}
			break
		}
	}

	return states
}
