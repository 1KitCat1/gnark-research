package circuits

import "github.com/consensys/gnark/frontend"

type SquareCircuit struct {
	X1 frontend.Variable
	X2 frontend.Variable
	A1 frontend.Variable `gnark:",public"`
	A2 frontend.Variable `gnark:",public"`
	A3 frontend.Variable `gnark:",public"`
}

func (circuit *SquareCircuit) VerifyRoot(X frontend.Variable, api frontend.API) {
	Xsq := api.Mul(X, X)
	result := api.Add(api.Mul(circuit.A1, Xsq),
		api.Mul(circuit.A2, X),
		circuit.A3)
	api.AssertIsEqual(result, 0)
}

func (circuit *SquareCircuit) Define(api frontend.API) error {
	circuit.VerifyRoot(circuit.X1, api)
	circuit.VerifyRoot(circuit.X2, api)

	return nil
}
