package main

import (
	"country-state-city-api/internal/adapters/handlers"
	"country-state-city-api/internal/adapters/persistence"
	"country-state-city-api/internal/ports/services"
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	_ = os.Setenv("HOST", "0.0.0.0")
	_ = os.Setenv("PORT", "8080")
	_ = os.Setenv("DATA_URL", "https://raw.githubusercontent.com/dr5hn/countries-states-cities-database/refs/heads/master/json/countries%2Bstates%2Bcities.json")
	_ = os.Setenv("SSL_ENABLED", "false")
	_ = os.Setenv("SSL_CERT", "/etc/letsencrypt/live/example.com/fullchain.pem")
	_ = os.Setenv("SSL_KEY", "/etc/letsencrypt/live/example.com/privkey.pem")

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
	router = gin.Default()
	router.Use(cors.Default())

	router.GET("/", handlers.HandleMain)
	router.GET("/countries", countryHandler.GetCountries)
	router.GET("/state/:country/names", stateHandler.GetStateNames)
	router.GET("/city/:country/:state/names", cityHandler.GetCityNames)

	log.Printf("Server running on http://localhost:%s\n", os.Getenv("PORT"))

	address := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	var err error
	if strings.EqualFold(os.Getenv("SSL_ENABLED"), "false") {
		err = router.Run(address)
	} else {
		err = router.Run(address, os.Getenv("SSL_CERT"), os.Getenv("SSL_KEY"))
	}

	if err != nil {
		log.Fatalf("Error starting the server: %s", err)
	}
}
