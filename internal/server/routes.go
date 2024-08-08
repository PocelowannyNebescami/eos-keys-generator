package server

import (
	"log"
	"net/http"

	"github.com/PocelowannyNebescami/eos-keys-generator/internal/keypair"
)

func (server *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// TODO: render index page
	http.Handle("GET /", http.FileServer(http.Dir("")))

	mux.HandleFunc("GET /key-pair", server.handleKeyPair)

	return mux
}

func (server *Server) handleKeyPair(w http.ResponseWriter, _ *http.Request) {
	_, err := keypair.NewRandomKeyPair()
	if err != nil {
		log.Println("Failed to generate a key pair: %w", err)
		http.Error(
			w,
			"Key pair was not generated",
			http.StatusInternalServerError,
		)

		return
	}

	// TODO: render template with a key pair
	http.Error(w, "", http.StatusNotImplemented)
}
