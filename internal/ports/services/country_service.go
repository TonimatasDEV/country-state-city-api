package services

import (
	"country-state-city-api/internal/ports/repositories"
)

type CountryService struct {
	countryRepo repositories.CountryRepository
}

func NewCountryService(countryRepo repositories.CountryRepository) *CountryService {
	return &CountryService{countryRepo: countryRepo}
}

func (service *CountryService) GetAllCountryNames() []string {
	countries := service.countryRepo.GetAll()

	var countryNames []string

	for _, country := range countries {
		countryNames = append(countryNames, country.Name)
	}

	return countryNames
}

func (service *CountryService) GetAllCountryIso2() []string {
	countries := service.countryRepo.GetAll()

	var countryNames []string

	for _, country := range countries {
		countryNames = append(countryNames, country.Iso2)
	}

	return countryNames
}
