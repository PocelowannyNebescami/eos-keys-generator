package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/PocelowannyNebescami/eos-keys-generator/cmd/web"
	"github.com/PocelowannyNebescami/eos-keys-generator/internal/keypair"
)

func (server *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.FileServerFS(web.Pages))

	mux.HandleFunc("GET /key-pair", server.handleKeyPair)

	return mux
}

func (server *Server) handleKeyPair(w http.ResponseWriter, _ *http.Request) {
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

	tmpl, err := template.New("keys").ParseFS(web.Pages, "*.html")
	if err != nil {
		log.Println("Failed to parse templates:", err)
		http.Error(
			w,
			"Failed to render the answer",
			http.StatusInternalServerError,
		)

		return
	}

	err = tmpl.ExecuteTemplate(w, "keys", keyPair)
	if err != nil {
		log.Println("Failed to execute template:", err)
		http.Error(
			w,
			"Failed to render the answer",
			http.StatusInternalServerError,
		)
	}
}
