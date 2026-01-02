package persistence

import (
	"country-state-city-api/internal/domain"
	"country-state-city-api/internal/ports/repositories"
	"encoding/json"
	"log"
	"os"
)

type JsonCountryRepository struct {
	data map[int]domain.Country
}

func (j JsonCountryRepository) Get() map[int]domain.Country {
	return j.data
}

func (j JsonCountryRepository) GetByID(id int) domain.Country {
	return j.data[id]
}

func NewJsonCountryRepository(jsonPath string) repositories.CountryRepository {
	file, err := os.Open(jsonPath)
	if err != nil {
		log.Fatalf("Error while opening file %v: %v", jsonPath, err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Error while closing file %v: %v", jsonPath, err)
		}
	}(file)

	decoder := json.NewDecoder(file)

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
