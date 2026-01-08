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
	_ = os.Setenv("PORT", "8080")
	_ = os.Setenv("DATA_URL", "https://raw.githubusercontent.com/dr5hn/countries-states-cities-database/refs/heads/master/json/countries%2Bstates%2Bcities.json")

	// Repositories
	countryRepo := persistence.NewJsonCountryRepository("countries+states+cities.json")

	// Services
	countryService := services.NewCountryService(countryRepo)
	stateService := services.NewStateService(countryRepo)
	cityService := services.NewCityService(countryRepo)

	// Handlers
	countryHandler := handlers.NewCountryHandler(countryService)
	stateHandler := handlers.NewStateHandler(stateService)
	cityHandler := handlers.NewCityHandler(cityService)

	// Router
	router := httprouter.New()

	router.GET("/", handlers.HandleMain)
	router.GET("/country/names", countryHandler.GetCountryNames)
	router.GET("/country/iso2", countryHandler.GetCountryIso2)
	router.GET("/state/:country/names", stateHandler.GetStateNames)
	router.GET("/city/:country/:state/names", cityHandler.GetCityNames)

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
