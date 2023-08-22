package main

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/pkg/errors"
	"gnark-research/circuits"
)

func main() {
	circuit := new(circuits.SquareCircuit)
	ecID := ecc.BN254

	// ---------------- COMPILE CIRCUIT -------------------
	ccs, err := frontend.Compile(ecID.ScalarField(), r1cs.NewBuilder, circuit)
	if err != nil {
		panic(errors.Wrap(err, "failed to compile circuit"))
	}

	// ---------------- SETUP PK AND VK --------------------
	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		panic(errors.Wrap(err, "failed to setup pk and vk"))
	}

	// ---------------- DEFINE WITNESS ----------
	assignment := &circuits.SquareCircuit{X1: -5, X2: 2, A1: 1, A2: 3, A3: -10}

	wit, err := frontend.NewWitness(assignment, ecID.ScalarField())
	if err != nil {
		panic(errors.Wrap(err, "failed to instantiate new witness"))
	}

	// ---------------- MAKE PROOF -------------------
	proof, err := groth16.Prove(ccs, pk, wit)
	if err != nil {
		panic(errors.Wrap(err, "failed to prove"))
	}

	// ----------------- VERIFY PROOF --------------------
	publicWitness, err := wit.Public()
	if err != nil {
		panic(errors.Wrap(err, "failed to extract public witness"))
	}
	if err := groth16.Verify(proof, vk, publicWitness); err != nil {
		panic(errors.Wrap(err, "failed to verify proof"))
	}

}
