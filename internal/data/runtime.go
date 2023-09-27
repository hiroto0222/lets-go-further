package data

import (
	"fmt"
	"strconv"
)

// Declare custom Runtime field for Movie.
type Runtime int32

// implements MarshalJSON to satisfy the json.Marshaler interface.
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	// Use the strconv.Quote() function on the string to wrap it in double quotes. It
	// needs to be surrounded by double quotes in order to be a valid *JSON string*
	quotedJsonValue := strconv.Quote(jsonValue)
	// Convert the quoted string value to a byte slice and return it.
	return []byte(quotedJsonValue), nil
}
