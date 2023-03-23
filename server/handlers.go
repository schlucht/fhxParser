package main

import "net/http"

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (app *application) AllUnits(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("load all units")
}
