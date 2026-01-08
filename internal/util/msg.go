package util

import (
	"encoding/json"
	"net/http"
)

type Msg struct {
	Msg string `json:"msg"`
}

type StringArrayJson struct {
	Array []string `json:"result"`
}

func SendError(w http.ResponseWriter, err error) {
	create(err.Error()).send(w)
}

func SendString(w http.ResponseWriter, status int, str string) {
	w.WriteHeader(status)
	create(str).send(w)
}

func SendStringArray(w http.ResponseWriter, status int, array []string) {
	SendJSON(w, status, StringArrayJson{array})
}

func create(str string) *Msg {
	return &Msg{Msg: str}
}

func (msg *Msg) send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(msg)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func SendJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(payload)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
