package handlers

import (
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

type SendCountries struct {
	Result []CountryJson `json:"result"`
}

type CountryJson struct {
	Name       string `json:"name"`
	NativeName string `json:"native_name"`
	Iso2       string `json:"iso2"`
	Iso3       string `json:"iso3"`
}

func (h *CountryHandler) GetCountries(c *gin.Context) {
	filters := util.GetFilters(c.Request)
	pop := filters["pop"]
	var filteredCountries []CountryJson

	countries := h.service.GetCountries()
	for i := 0; i < len(countries); i++ {
		country := countries[i]

		countryJson := CountryJson{
			Name:       country.Name,
			NativeName: country.Native,
			Iso2:       country.Iso2,
			Iso3:       country.Iso3,
		}

		if !pop.IsEmpty() {
			if country.Population > pop.Int() {
				filteredCountries = append(filteredCountries, countryJson)
			}
		} else {
			filteredCountries = append(filteredCountries, countryJson)
		}
	}

	result := SendCountries{
		Result: filteredCountries,
	}

	c.JSON(http.StatusOK, result)
}
