package handlers

import (
	"country-state-city-api/internal/ports/services"
	"country-state-city-api/internal/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type StateHandler struct {
	service *services.StateService
}

func NewStateHandler(service *services.StateService) *StateHandler {
	return &StateHandler{service: service}
}

func (h *StateHandler) GetStateNames(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	util.SendJSON(w, http.StatusOK, h.service.GetStatesNames(p.ByName("country")))
}
