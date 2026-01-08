package handlers

import (
	"country-state-city-api/internal/ports/services"
	"country-state-city-api/internal/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CountryHandler struct {
	service *services.CountryService
}

func NewCountryHandler(service *services.CountryService) *CountryHandler {
	return &CountryHandler{service: service}
}

func (h *CountryHandler) GetCountryNames(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	filters := util.GetFilters(r)
	pop := filters["pop"]
	var filteredCountries []string

	for _, country := range h.service.GetCountries() {
		if !pop.IsEmpty() {
			if country.Population > pop.Int() {
				filteredCountries = append(filteredCountries, country.Name)
			}
		} else {
			filteredCountries = append(filteredCountries, country.Name)
		}
	}

	util.SendStringArray(w, http.StatusOK, filteredCountries)
}

func (h *CountryHandler) GetCountryIso2(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	util.SendStringArray(w, http.StatusOK, h.service.GetAllCountryIso2())
}
