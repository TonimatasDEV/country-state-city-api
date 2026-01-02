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

func (h *CountryHandler) GetCountryNames(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	util.SendJSON(w, http.StatusOK, h.service.GetAllCountryNames())
}

func (h *CountryHandler) GetCountryIso2(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	util.SendJSON(w, http.StatusOK, h.service.GetAllCountryIso2())
}
