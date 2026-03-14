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

func (service *CountryService) GetCountries() []domain.Country {
	return service.countryRepo.Get()
}
