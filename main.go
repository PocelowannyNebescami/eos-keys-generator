package main

import (
	"fmt"
	"github.com/eoscanada/eos-go/ecc"
)

func main() {
	pvtKey, err := ecc.NewRandomPrivateKey()
	if err != nil {
		fmt.Println("Failed to create a key")
		return
	}

	fmt.Println("Private key:", pvtKey.String())
	fmt.Println("Public key:", pvtKey.PublicKey().String())
}
