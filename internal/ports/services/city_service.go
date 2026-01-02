package services

import (
	"country-state-city-api/internal/ports/repositories"
	"strings"
)

type CityService struct {
	countryRepo repositories.CountryRepository
}

func NewCityService(countryRepo repositories.CountryRepository) *CityService {
	return &CityService{countryRepo: countryRepo}
}

func (service *CityService) GetCityNames(countryName string, stateName string) []string {
	countries := service.countryRepo.Get()

	var cities []string

	for _, country := range countries {
		if strings.EqualFold(countryName, country.Name) {
			for _, state := range country.States {
				if strings.EqualFold(stateName, state.Name) {
					for _, city := range state.Cities {
						cities = append(cities, city.Name)
					}

					if cities == nil {
						cities = append(cities, stateName)
					}

					break
				}
			}
		}
	}

	return cities
}
