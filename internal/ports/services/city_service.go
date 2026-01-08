package services

import (
	"country-state-city-api/internal/ports/repositories"
)

type CityService struct {
	countryRepo repositories.CountryRepository
}

func NewCityService(countryRepo repositories.CountryRepository) *CityService {
	return &CityService{countryRepo: countryRepo}
}

func (service *CityService) GetCityNames(countryName string, stateName string) []string {
	cities := service.countryRepo.GetCities(countryName, stateName)

	var cityNames []string

	for _, city := range cities {
		cityNames = append(cityNames, city.Name)
	}

	if len(cityNames) == 0 {
		return []string{stateName}
	}

	return cityNames
}
