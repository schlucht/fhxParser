package main

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (app *application) RegisterPage(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "register", &templateData{}); err != nil {
		app.errorLog.Println(err)
	}
}

// User in der Datenbank speichern
func (app *application) SaveNewUser(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Daten aus dem Body holen
	err := app.readJSON(w, r, &userInput)
	if err != nil {
		app.badRequest(w, r, err, "SaveNewUser: Decode")
		return
	}

	// Passwort verschluesseln
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), 12)
	if err != nil {
		app.badRequest(w, r, err, "SaveNewUser: GenerateFromPassword")
		return
	}

	// Daten in die Datenbank schreiben
	err = app.DB.CreateNewUser(userInput.Name, userInput.Email, string(hashedPassword))
	if err != nil {
		app.badRequest(w, r, err, "SaveNewUser: CreateNewUser")
		return
	}

	j := jsonResponse{
		OK:      true,
		Message: "User saved to database",
		Content: "",
		ID:      0,
	}
	err = app.writeJSON(w, http.StatusOK, j)
	if err != nil {
		app.badRequest(w, r, err, "SaveNewUser: writeJSON")
		return
	}
	// app.Session.Put(r.Context(), "userID", usr.ID)

	// http.Redirect(w, r, "/login", http.StatusSeeOther)
}
