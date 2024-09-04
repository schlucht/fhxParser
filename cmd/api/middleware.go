package main

import (
	"net/http"
)

// Wenn keine Anlagen gespeichert sind, wird beim Einlesen einer
// fhx Datei auf die Anlageseite weitergeleitet.
func (app *application) Plant(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		plants, err := app.LoadPlants()
		if err != nil {
			app.errorLog.Println(err)
		}
		if len(plants) == 0 {
			http.Redirect(w, r, "/plants", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
