package main

import (
	"fmt"
	"saifas.org/eos-key-generator/keypair"
)

func main() {
	keyPair, err := keypair.NewRandomKeyPair()
	if err != nil {
		fmt.Println("Failed to generate a key pair: %w", err)
	}

	fmt.Println("Private key:", keyPair.Pvt)
	fmt.Println("Public key:", keyPair.Pub)
}
