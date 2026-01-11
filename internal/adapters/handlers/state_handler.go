package handlers

import (
	"country-state-city-api/internal/domain"
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

func (h *StateHandler) GetStateNames(c *gin.Context) {
	result := domain.StringArrayJson{
		Array: h.service.GetStatesNames(c.Params.ByName("country")),
	}

	c.JSON(http.StatusOK, result)
}
