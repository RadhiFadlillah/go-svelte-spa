package backend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) getHello(w http.ResponseWriter, r *http.Request) {
	// If error ever occured, return HTTP error
	var err error
	defer func() { markHttpError(w, err) }()

	// Prepare message
	vars := mux.Vars(r)
	msg := struct{ Message string }{
		Message: fmt.Sprintf("Hello %s", vars["name"]),
	}

	// Encode to JSON
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&msg)
}
