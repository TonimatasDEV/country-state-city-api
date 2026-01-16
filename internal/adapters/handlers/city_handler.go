package handlers

import (
	"country-state-city-api/internal/domain"
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

func (h *CityHandler) GetCities(c *gin.Context) {
	result := domain.StringArrayJson{
		Array: h.service.GetCityNames(c.Params.ByName("country"), c.Params.ByName("state")),
	}

	c.JSON(http.StatusOK, result)
}
