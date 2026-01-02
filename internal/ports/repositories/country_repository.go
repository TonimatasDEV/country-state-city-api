package repositories

import "country-state-city-api/internal/domain"

type CountryRepository interface {
	GetAll() []domain.Country
	GetByID(id int) domain.Country
}
