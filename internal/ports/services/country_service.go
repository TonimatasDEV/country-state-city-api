package services

import (
	"country-state-city-api/internal/ports/repositories"
	"log"
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
		log.Println(country.Name)
		countryNames = append(countryNames, country.Name)
	}

	return countryNames
}
