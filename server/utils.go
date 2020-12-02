package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SetError writes an error response to a client.
func SetError(w http.ResponseWriter, err error) {
	fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
}

// SetJSONResponse responds using a JSON object.
func SetJSONResponse(w http.ResponseWriter, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(b))
}
