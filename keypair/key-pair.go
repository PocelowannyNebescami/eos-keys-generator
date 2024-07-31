package keypair

import (
	"github.com/eoscanada/eos-go/ecc"
)

type KeyPair struct {
	Pub string `json:"pub"`
	Pvt string `json:"pvt"`
}

func NewRandomKeyPair() (KeyPair, error) {
	pvtKey, err := ecc.NewRandomPrivateKey()
	if err != nil {
		return KeyPair{}, err
	}

	return KeyPair{
		Pub: pvtKey.PublicKey().String(),
		Pvt: pvtKey.String(),
	}, nil
}
