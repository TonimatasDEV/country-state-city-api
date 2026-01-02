package handlers

import (
	"country-state-city-api/internal/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HandleMain(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	util.SendString(w, http.StatusOK, "Hello world! This is the best api to get countries, states and cities!")
}
