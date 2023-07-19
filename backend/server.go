package backend

import (
	"database/sql"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

// Server handles routing for the app
type Server struct {
	Assets fs.FS
}

// Serve serves app in specified port.
func (s *Server) Serve(port int) error {
	// Create router
	router := mux.NewRouter()

	// Handle API first
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/hello/{name}", s.getHello).Methods("GET")

	// Handle UI route later
	uiRouter := router.PathPrefix("/").Subrouter()
	uiRouter.Methods("GET").HandlerFunc(s.serveAssets)

	// Create server
	url := fmt.Sprintf(":%d", port)
	svr := &http.Server{
		Addr:         url,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: time.Minute,
	}

	// Serve app
	log.Printf("serve app in %s", url)
	return svr.ListenAndServe()
}

func markHttpError(w http.ResponseWriter, err error) {
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, err.Error(), 500)
	}
}
