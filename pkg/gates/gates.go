package gates

import (
	"math"
	"math/cmplx"
	"testing"

	"quantum-calculations/pkg/qubits"
)

func X() [][]complex128 {
	return [][]complex128{
		{0, 1},
		{1, 0},
	}
}

func Y() [][]complex128 {
	return [][]complex128{
		{0, -1i},
		{1i, -1},
	}
}

func Z() [][]complex128 {
	return [][]complex128{
		{1, 0},
		{0, -1},
	}
}

func Hadamard() [][]complex128 {
	sqrt2 := cmplx.Sqrt(complex(2, 0))
	return [][]complex128{
		{1 / sqrt2, 1 / sqrt2},
		{1 / sqrt2, -1 / sqrt2},
	}
}

func CNOT() [][]complex128 {
	return [][]complex128{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func S() [][]complex128 {
	return [][]complex128{
		{1, 0},
		{0, 1i},
	}
}

func T() [][]complex128 {
	phase := cmplx.Exp(1i * math.Pi / 4) // e^(i\pi/4)
	return [][]complex128{
		{1, 0},
		{0, phase},
	}
}

func TestTGate(t *testing.T) {
	state := []complex128{0, 1}
	qubit := qubits.NewQubit(state)

	qubit.ApplyGate(T())

	expected := []complex128{0, cmplx.Exp(1i * math.Pi / 4)}
	for i := range qubit.State {
		if qubit.State[i] != expected[i] {
			t.Errorf("Estado incorreto após aplicar a porta T: esperado %v, obtido %v", expected, qubit.State)
		}
	}

}
