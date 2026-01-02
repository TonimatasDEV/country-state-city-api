package repositories

import "country-state-city-api/internal/domain"

type CountryRepository interface {
	Get() map[int]domain.Country
	GetByID(id int) domain.Country
}
