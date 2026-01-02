package repositories

import "contry-state-city-api/internal/domain"

type CountryRepository interface {
	GetAll() ([]domain.Country, error)
	GetByID(id int) (domain.Country, error)
}
