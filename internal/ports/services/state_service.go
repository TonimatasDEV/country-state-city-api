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
	states := service.countryRepo.GetStates(countryCode)

	// Start - Temporal fix to Spanish states
	var result []domain.State

	if countryCode == "es" {
		for _, state := range states {
			if state.Type == "autonomous community" {
				continue
			}

			result = append(result, state)
		}
	} else {
		result = append(result, states...)
	}
	// End

	return result
}
