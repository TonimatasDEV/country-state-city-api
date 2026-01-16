package services

import (
	"country-state-city-api/internal/domain"
	"country-state-city-api/internal/ports/repositories"
)

type CityService struct {
	countryRepo repositories.CountryRepository
}

func NewCityService(countryRepo repositories.CountryRepository) *CityService {
	return &CityService{countryRepo: countryRepo}
}

func (service *CityService) GetCities(countryCode string, stateCode string) []domain.City {
	return service.countryRepo.GetCities(countryCode, stateCode)
}
