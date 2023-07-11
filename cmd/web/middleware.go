package main

import "net/http"

func (app *application) LoadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
