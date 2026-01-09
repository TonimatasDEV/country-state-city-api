package handlers

import (
	"country-state-city-api/internal/domain"
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
	h.sendCountries(w, r, func(country domain.Country) string {
		return country.Name
	})
}

func (h *CountryHandler) GetCountryNativeNames(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.sendCountries(w, r, func(country domain.Country) string {
		return country.Native
	})
}

func (h *CountryHandler) GetCountryIso2(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.sendCountries(w, r, func(country domain.Country) string {
		return country.Iso2
	})
}

func (h *CountryHandler) GetCountryIso3(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.sendCountries(w, r, func(country domain.Country) string {
		return country.Iso3
	})
}

func (h *CountryHandler) sendCountries(w http.ResponseWriter, r *http.Request, getInfo func(country domain.Country) string) {
	filters := util.GetFilters(r)
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

	util.SendStringArray(w, http.StatusOK, filteredCountries)
}
