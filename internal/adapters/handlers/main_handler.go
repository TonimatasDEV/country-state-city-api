package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainMessage struct {
	Msg          string `json:"msg"`
	Wiki         string `json:"wiki"`
	License      string `json:"license"`
	Repository   string `json:"repository"`
	IssueTracker string `json:"issueTracker"`
}

func HandleMain(c *gin.Context) {
	mainMessage := MainMessage{
		Msg:          "Hello World! This is the best api to get countries, states and cities!",
		Wiki:         "https://github.com/TonimatasDEV/country-state-city-api/wiki",
		License:      "https://github.com/TonimatasDEV/country-state-city-api/blob/master/LICENSE",
		Repository:   "https://github.com/TonimatasDEV/country-state-city-api",
		IssueTracker: "https://github.com/TonimatasDEV/country-state-city-api/issues",
	}

	c.JSON(http.StatusOK, mainMessage)
}
