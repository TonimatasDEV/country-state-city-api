package handlers

import (
	"country-state-city-api/internal/ports/services"
	"country-state-city-api/internal/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CityHandler struct {
	service *services.CityService
}

func NewCityHandler(service *services.CityService) *CityHandler {
	return &CityHandler{service: service}
}

func (h *CityHandler) GetCityNames(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	util.SendStringArray(w, http.StatusOK, h.service.GetCityNames(p.ByName("country"), p.ByName("state")))
}
