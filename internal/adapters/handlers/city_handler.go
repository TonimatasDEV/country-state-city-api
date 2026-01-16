package handlers

import (
	"country-state-city-api/internal/ports/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CityHandler struct {
	service *services.CityService
}

func NewCityHandler(service *services.CityService) *CityHandler {
	return &CityHandler{service: service}
}

type SendCities struct {
	Result []CityJson `json:"result"`
}

type CityJson struct {
	Name string `json:"name"`
}

func (h *CityHandler) GetCities(c *gin.Context) {
	var cities []CityJson

	for _, city := range h.service.GetCities(c.Params.ByName("country"), c.Params.ByName("state")) {
		cities = append(cities, CityJson{city.Name})
	}

	result := SendCities{
		Result: cities,
	}

	c.JSON(http.StatusOK, result)
}
