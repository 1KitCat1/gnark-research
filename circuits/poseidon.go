package circuits

import (
	"gnark-research/poseidon"
	"math/big"

	"github.com/consensys/gnark/frontend"
)

type Poseidon struct {
	Input frontend.Variable `gnark:"input"`
	Hash  frontend.Variable `gnark:",public"`
}

func (circuit *Poseidon) Define(api frontend.API) error {
	poseidonHash := poseidon.NewPoseidon1(api)
	poseidonHash.Write(circuit.Input)
	api.AssertIsEqual(circuit.Hash, poseidonHash.Sum())
	return nil
}

func NewPoseidonAssignment() *Poseidon {
	hashValue, _ := new(big.Int).SetString("13377623690824916797327209540443066247715962236839283896963055328700043345550", 0)
	return &Poseidon{
		Input: 111,
		Hash:  hashValue,
	}
}
