package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/PocelowannyNebescami/eos-keys-generator/internal/keypair"
)

func main() {
	// TODO: fix templates parsing
	allTemplates, err := template.New("all").ParseGlob("*.html")
	if err != nil {
		panic("Failed to parse templates: " + err.Error())
	}

	http.HandleFunc(
		"GET /key-pair",
		func(w http.ResponseWriter, _ *http.Request) {
			keyPair, err := keypair.NewRandomKeyPair()
			if err != nil {
				log.Println("Failed to generate a key pair: %w", err)
				http.Error(
					w,
					"Key pair was not generated",
					http.StatusInternalServerError,
				)

				return
			}

			err = allTemplates.ExecuteTemplate(w, "keys", keyPair)
			if err != nil {
				log.Println(err)
				http.Error(w, "", http.StatusInternalServerError)
			}
		},
	)

	// TODO: fix
	http.Handle("GET /", http.FileServer(http.Dir("./views/dist")))

	log.Fatalln(http.ListenAndServe(":9090", nil))
}