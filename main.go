package main

import (
	"html/template"
	"log"
	"net/http"

	"saifas.org/eos-key-generator/keypair"
)

const indexTemplatePath string = "./views/index.html"

func main() {
	index, err := template.New("index").ParseFiles(indexTemplatePath)
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

			err = index.ExecuteTemplate(w, "keys", keyPair)
			if err != nil {
				log.Println(err)
				http.Error(w, "", http.StatusInternalServerError)
			}
		},
	)

	http.HandleFunc(
		"GET /",
		func(w http.ResponseWriter, _ *http.Request) {
			err = index.ExecuteTemplate(w, "index", nil)
			if err != nil {
				log.Println(err)
				http.Error(w, "", http.StatusInternalServerError)
			}
		},
	)

	log.Fatalln(http.ListenAndServe(":9090", nil))
}
