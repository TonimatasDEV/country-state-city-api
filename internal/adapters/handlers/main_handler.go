package handlers

import (
	"country-state-city-api/internal/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MainMessage struct {
	Msg          string `json:"msg"`
	Wiki         string `json:"wiki"`
	License      string `json:"license"`
	Repository   string `json:"repository"`
	IssueTracker string `json:"issueTracker"`
}

func HandleMain(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	mainMessage := MainMessage{
		Msg:          "Hello World! This is the best api to get countries, states and cities!",
		Wiki:         "https://github.com/TonimatasDEV/country-state-city-api/wiki",
		License:      "https://github.com/TonimatasDEV/country-state-city-api/blob/master/LICENSE",
		Repository:   "https://github.com/TonimatasDEV/country-state-city-api",
		IssueTracker: "https://github.com/TonimatasDEV/country-state-city-api/issues",
	}

	util.SendJSON(w, http.StatusOK, mainMessage)
}
