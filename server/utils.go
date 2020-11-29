package server

import (
	"fmt"
	"net/http"
)

// SetError writes an error response to a client.
func SetError(w http.ResponseWriter, err error) {
	fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
}
