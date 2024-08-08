package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PocelowannyNebescami/eos-keys-generator/internal/server"
)

func main() {
	server := server.NewServer()
	err := server.ListenAndServe()

	if err != http.ErrServerClosed {
		log.Fatal(err)
		os.Exit(1)
	}
}
