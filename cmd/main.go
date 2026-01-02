package main

import (
	"country-state-city-api/internal/adapters/handlers"
	"country-state-city-api/internal/adapters/persistence"
	"country-state-city-api/internal/ports/services"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	os.Setenv("PORT", "8080")

	// Repositories
	countryRepo := persistence.NewJsonCountryRepository("countries+states+cities.json")

	// Services
	countryService := services.NewCountryService(countryRepo)
	stateService := services.NewStateService(countryRepo)

	// Handlers
	countryHandler := handlers.NewCountryHandler(countryService)
	stateHandler := handlers.NewStateHandler(stateService)

	// Router
	router := httprouter.New()

	router.GET("/", handlers.HandleMain)
	router.GET("/country/names", countryHandler.GetCountriesByName)
	router.GET("/country/iso2", countryHandler.GetCountriesByIso2)
	router.GET("/state/:country/names", stateHandler.GetStateNames)

	log.Printf("Server running on http://localhost:%s\n", os.Getenv("PORT"))

	server := &http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Error starting the server: %s", err)
	}
}
