package gates

import (
	"math"
	"math/cmplx"
	"testing"

	"github.com/SuperpositionLabs/quantum-calculations/pkg/qubits"
)

// X returns the Pauli-X (NOT) gate matrix.
func X() [][]complex128 {
	return [][]complex128{
		{0, 1},
		{1, 0},
	}
}

// Y returns the Pauli-Y gate matrix.
func Y() [][]complex128 {
	return [][]complex128{
		{0, -1i},
		{1i, -1},
	}
}

// Z returns the Pauli-Z gate matrix.
func Z() [][]complex128 {
	return [][]complex128{
		{1, 0},
		{0, -1},
	}
}

// Hadamard returns the Hadamard gate matrix.
func Hadamard() [][]complex128 {
	sqrt2 := cmplx.Sqrt(complex(2, 0))
	return [][]complex128{
		{1 / sqrt2, 1 / sqrt2},
		{1 / sqrt2, -1 / sqrt2},
	}
}

// CNOT returns the Controlled-NOT (CNOT) gate matrix.
func CNOT() [][]complex128 {
	return [][]complex128{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

// S returns the S-phase gate matrix.
func S() [][]complex128 {
	return [][]complex128{
		{1, 0},
		{0, 1i},
	}
}

// T returns the T-phase gate matrix.
func T() [][]complex128 {
	phase := cmplx.Exp(1i * math.Pi / 4) // e^(iπ/4)
	return [][]complex128{
		{1, 0},
		{0, phase},
	}
}

// TestTGate verifies the correctness of the T gate operation.
func TestTGate(t *testing.T) {
	state := []complex128{0, 1}
	qubit := qubits.NewQubit(state)

	qubit.ApplyGate(T())

	expected := []complex128{0, cmplx.Exp(1i * math.Pi / 4)}
	for i := range qubit.State {
		if qubit.State[i] != expected[i] {
			t.Errorf("Incorrect state after applying the T gate: expected %v, got %v", expected, qubit.State)
		}
	}
}
