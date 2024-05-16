package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/schlucht/fhxreader/internal/models"
)

func (app *application) LoginPage(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "login", &templateData{}); err != nil {
		app.errorLog.Println(err)
	}
}
func (app *application) CreateAuthToken(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &userInput)
	if err != nil {
		app.badRequest(w, r, err, "CreateAuthToken: readJson")
		return
	}

	// User mit der Email holen
	usr, err := app.DB.GetUserByEmail(userInput.Email)
	if err != nil {
		app.invalidCredentials(w)
		return
	}

	// Passwort vergleichen
	valid, err := app.passwordMatches(usr.Password, userInput.Password)
	app.infoLog.Println(userInput)
	if (err != nil) || !valid {
		app.invalidCredentials(w)
		return
	}
	if !valid {
		app.invalidCredentials(w)
		return
	}

	// Token erstellen
	token, err := models.GenerateToken(usr.ID, 24*time.Hour, models.ScopeAuthentication)
	if err != nil {
		app.badRequest(w, r, err, "CreateAuthToken: GenerateToken")
		return
	}

	// Token speichern
	err = app.DB.InsertToken(token, usr)
	if err != nil {
		app.badRequest(w, r, err, "CreateAuthToken: InsertToken")
		return
	}

	j := jsonResponse{
		OK:      true,
		Message: fmt.Sprintf("token for %s created", userInput.Email),
		Content: token,
		ID:      0,
	}
	err = app.writeJSON(w, http.StatusOK, j)
	if err != nil {
		app.badRequest(w, r, err, "CreateAuthToken: writeJSON")
		return
	}
}
