package handlers

import (
	"country-state-city-api/internal/domain"
	"country-state-city-api/internal/ports/services"
	"country-state-city-api/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	service *services.CountryService
}

func NewCountryHandler(service *services.CountryService) *CountryHandler {
	return &CountryHandler{service: service}
}

func (h *CountryHandler) GetCountryNames(c *gin.Context) {
	h.sendCountries(c, func(country domain.Country) string {
		return country.Name
	})
}

func (h *CountryHandler) GetCountryNativeNames(c *gin.Context) {
	h.sendCountries(c, func(country domain.Country) string {
		return country.Native
	})
}

func (h *CountryHandler) GetCountryIso2(c *gin.Context) {
	h.sendCountries(c, func(country domain.Country) string {
		return country.Iso2
	})
}

func (h *CountryHandler) GetCountryIso3(c *gin.Context) {
	h.sendCountries(c, func(country domain.Country) string {
		return country.Iso3
	})
}

func (h *CountryHandler) sendCountries(c *gin.Context, getInfo func(country domain.Country) string) {
	filters := util.GetFilters(c.Request)
	pop := filters["pop"]
	var filteredCountries []string

	for _, country := range h.service.GetCountries() {
		if !pop.IsEmpty() {
			if country.Population > pop.Int() {
				filteredCountries = append(filteredCountries, getInfo(country))
			}
		} else {
			filteredCountries = append(filteredCountries, getInfo(country))
		}
	}

	result := domain.StringArrayJson{
		Array: filteredCountries,
	}

	c.JSON(http.StatusOK, result)
}
