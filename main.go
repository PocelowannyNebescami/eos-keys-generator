package main

import (
	"log"
	"net/http"
	"encoding/json"

	"saifas.org/eos-key-generator/keypair"
)

func main() {
	http.HandleFunc(
		"GET /key-pair",
		func(w http.ResponseWriter, _ *http.Request) {
			keyPair, err := keypair.NewRandomKeyPair()
			if err != nil {
				log.Println("Failed to generate a key pair: %w", err)
				return // TODO: add status code
			}

			jsonResponse, err := json.Marshal(keyPair)
			if err != nil {
				log.Println("Failed to convert key pair to JSON: %w", err)
				return // TODO: add status code
			}

			_, err = w.Write(jsonResponse)
			if err != nil {
				log.Println(err)
			}
		},
	)

	log.Println(http.ListenAndServe(":9090", nil))
}
