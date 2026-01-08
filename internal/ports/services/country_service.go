package services

import (
	"country-state-city-api/internal/domain"
	"country-state-city-api/internal/ports/repositories"
)

type CountryService struct {
	countryRepo repositories.CountryRepository
}

func NewCountryService(countryRepo repositories.CountryRepository) *CountryService {
	return &CountryService{countryRepo: countryRepo}
}

func (service *CountryService) GetCountries() map[int]domain.Country {
	return service.countryRepo.Get()
}

func (service *CountryService) GetAllCountryIso2() []string {
	countries := service.countryRepo.Get()

	var countryNames []string

	for _, country := range countries {
		countryNames = append(countryNames, country.Iso2)
	}

	return countryNames
}
