package main

import (
	"go-svelte-spa/backend"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	// Prepare logger
	zerolog.TimeFieldFormat = time.DateTime

	// Start server
	srv := backend.Server{Assets: assets}
	srv.Serve(8080)
}
