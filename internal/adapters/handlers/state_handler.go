package handlers

import (
	"country-state-city-api/internal/ports/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StateHandler struct {
	service *services.StateService
}

func NewStateHandler(service *services.StateService) *StateHandler {
	return &StateHandler{service: service}
}

type SendStates struct {
	Result []StateJson `json:"result"`
}

type StateJson struct {
	Name       string `json:"name"`
	NativeName string `json:"native_name"`
	Iso2       string `json:"iso2"`
}

func (h *StateHandler) GetStates(c *gin.Context) {
	var states []StateJson

	for _, state := range h.service.GetStates(c.Params.ByName("country")) {
		states = append(states, StateJson{
			Name:       state.Name,
			NativeName: state.Native,
			Iso2:       state.Iso2,
		})
	}

	result := SendStates{
		Result: states,
	}

	c.JSON(http.StatusOK, result)
}
