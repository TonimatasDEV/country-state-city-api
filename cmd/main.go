package main

import (
	"contry-state-city-api/internal/adapters/persistence"
	"log"
	"time"
)

func main() {
	// Repositories
	countryRepo := persistence.NewJsonCountryRepository("countries+states+cities.json")

	// Services

	// Handlers

	start := time.Now().UnixMilli()
	log.Print(countryRepo.GetAll())
	elapsed := time.Now().UnixMilli() - start
	log.Printf("%dms", elapsed)
}
