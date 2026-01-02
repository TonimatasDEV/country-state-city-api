package persistence

import (
	"contry-state-city-api/internal/domain"
	"contry-state-city-api/internal/ports/repositories"
	json2 "encoding/json"
	"log"
	"os"
)

type JsonCountryRepository struct {
	data map[int]domain.Country
}

func (j JsonCountryRepository) GetAll() ([]domain.Country, error) {
	values := make([]domain.Country, len(j.data))

	for _, country := range j.data {
		values = append(values, country)
	}

	return values, nil
}

func (j JsonCountryRepository) GetByID(id int) (domain.Country, error) {
	return j.data[id], nil
}

func NewJsonCountryRepository(jsonPath string) repositories.CountryRepository {
	file, err := os.Open(jsonPath)
	if err != nil {
		log.Fatalf("Error while opening file %v: %v", jsonPath, err)
	}

	defer file.Close()

	decoder := json2.NewDecoder(file)

	var countries []domain.Country

	if err := decoder.Decode(&countries); err != nil {
		log.Fatalf("Error while decoding file %v: %v", jsonPath, err)
	}

	var countryMap = make(map[int]domain.Country)

	for _, country := range countries {
		countryMap[country.ID] = country
	}

	return JsonCountryRepository{data: countryMap}
}
