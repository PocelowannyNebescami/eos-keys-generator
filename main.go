package main

import (
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
			}

			_, err = w.Write([]byte(keyPair.Pub + " " + keyPair.Pvt))
			if err != nil {
				log.Println(err)
			}
		},
	)

	log.Println(http.ListenAndServe(":9090", nil))
}
