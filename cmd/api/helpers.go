package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

// writeJSON writes aribtrary data out as JSON
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

// readJSON reads json from request body into data. We only accept a single json value in the body
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576 // max one megabyte in request body
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	// we only allow one entry in the json file
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only have a single JSON value")
	}

	return nil
}

// badRequest sends a JSON response with status http.StatusBadRequest, describing the error
func (app *application) badRequest(w http.ResponseWriter, err error, method string, statuscode ...int) error {	
	status := http.StatusBadRequest
	if len(statuscode) == 0 {
		status = statuscode[0]
	}

	var payload jsonResponse
	var customErr error

	switch {
	case strings.Contains(err.Error(), "Error 1062"):
		customErr = errors.New("duplicate value violates unique constrain")
		status = http.StatusForbidden
	case strings.Contains(err.Error(), "SQLSTATE 22001"):
		customErr = errors.New("the value you are trying to insert is to large")
		status = http.StatusForbidden
	case strings.Contains(err.Error(), "SQLSTATE 23403"):
		customErr = errors.New("foreign key violation")
		status = http.StatusForbidden
	default:
		customErr = err
	}

	payload.Error = true
	payload.Message = customErr.Error()

	if err = app.writeJSON(w, status, payload); err != nil {
		return err
	}
	return nil
}
