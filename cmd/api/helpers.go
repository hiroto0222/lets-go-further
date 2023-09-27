package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]any

// readIDParam retrieves the "id" URL parameter from the current request context, then converts it to
// an integer and returns it. If the operation isn't successful, returns 0 and an error.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	// Get URL parameters stored in the request context.
	params := httprouter.ParamsFromContext(r.Context())

	// Get ID from param.
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

// writeJSON takes the http.ResponseWriter, the HTTP status code, and the data to encode to JSON,
// and a header map containing any additional HTTP headers to include in the response.
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// Encode the data to JSON
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')
	// Add headers
	for k, v := range headers {
		w.Header()[k] = v
	}

	// Add content-type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
