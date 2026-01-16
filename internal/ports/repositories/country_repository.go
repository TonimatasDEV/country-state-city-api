package repositories

import "country-state-city-api/internal/domain"

type CountryRepository interface {
	Get() map[int]domain.Country
	GetStates(countryCode string) []domain.State
	GetCities(countryCode, stateCode string) []domain.City
	GetByID(id int) domain.Country
}
