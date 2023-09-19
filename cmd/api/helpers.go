package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

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
