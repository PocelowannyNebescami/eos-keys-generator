package main

import (
	"encoding/json"
	"log"
	"net/http"

	"saifas.org/eos-key-generator/keypair"
)

func main() {
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

			jsonResponse, err := json.Marshal(keyPair)
			if err != nil {
				log.Println("Failed to convert key pair to JSON: %w", err)
				http.Error(
					w,
					"Response was not formed correctly",
					http.StatusInternalServerError,
				)

				return
			}

			_, err = w.Write(jsonResponse)
			if err != nil {
				log.Println(err)
				http.Error(w, "", http.StatusInternalServerError)
			}
		},
	)

	log.Fatalln(http.ListenAndServe(":9090", nil))
}
