package persistence

import (
	"country-state-city-api/internal/domain"
	"country-state-city-api/internal/ports/repositories"
	"country-state-city-api/internal/util"
	"encoding/json"
	"errors"
	"io/fs"
	"log"
	"os"
	"strings"
)

type JsonCountryRepository struct {
	data map[int]domain.Country
}

func (j JsonCountryRepository) Get() map[int]domain.Country {
	return j.data
}

func (j JsonCountryRepository) GetStates(countryCode string) []domain.State {
	for _, country := range j.data {
		if strings.EqualFold(country.Iso2, countryCode) {
			return country.States
		}
	}

	return []domain.State{}
}

func (j JsonCountryRepository) GetCities(countryName, stateName string) []domain.City {
	for _, state := range j.GetStates(countryName) {
		if strings.EqualFold(state.Name, stateName) {
			return state.Cities
		}
	}

	return []domain.City{}
}

func (j JsonCountryRepository) GetByID(id int) domain.Country {
	return j.data[id]
}

func NewJsonCountryRepository(jsonPath string) repositories.CountryRepository {
	file, err := os.Open(jsonPath)

	if errors.Is(err, fs.ErrNotExist) {
		log.Println("Downloading data...")
		err = util.DownloadFile(jsonPath, os.Getenv("DATA_URL"))

		if err != nil {
			log.Println("Error downloading data.")
		} else {
			log.Println("Data downloaded correctly.")
			return NewJsonCountryRepository(jsonPath)
		}
	}

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
