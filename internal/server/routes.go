package server

import (
	"log"
	"net/http"

	"github.com/PocelowannyNebescami/eos-keys-generator/cmd/web"
	"github.com/PocelowannyNebescami/eos-keys-generator/internal/keypair"
	"github.com/a-h/templ"
)

func (server *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /", templ.Handler(web.Index()))

	mux.HandleFunc("GET /key-pair", server.handleKeyPair)

	mux.Handle("GET /assets/", http.FileServerFS(web.Assets))

	return mux
}

func (server *Server) handleKeyPair(w http.ResponseWriter, r *http.Request) {
	keyPair, err := keypair.NewRandomKeyPair()
	if err != nil {
		log.Println("Failed to generate a key pair:", err)
		http.Error(
			w,
			"Key pair was not generated",
			http.StatusInternalServerError,
		)

		return
	}

	keyComponent := web.Keys(keyPair)
	err = keyComponent.Render(r.Context(), w)
	if err != nil {
		log.Println("Failed to render the template:", err)
		http.Error(
			w,
			"Failed to render the answer",
			http.StatusInternalServerError,
		)
	}
}
