package qubits

import (
	"fmt"
	"math/cmplx"
	"math/rand"
)

// Qubit represents a qubit with its state as a slice of complex numbers.
type Qubit struct {
	State []complex128
}

// NewQubit returns a new normalized qubit from the given state.
func NewQubit(state []complex128) *Qubit {
	q := &Qubit{State: state}
	q.Normalize()
	return q
}

// NewNNQubit returns a new qubit without normalizing its initial state.
func NewNNQubit(state []complex128) *Qubit {
	return &Qubit{State: state}
}

// Normalize normalizes the qubit's state to ensure the sum of probabilities is 1.
func (q *Qubit) Normalize() {
	norm := 0.0
	for _, amp := range q.State {
		norm += cmplx.Abs(amp) * cmplx.Abs(amp)
	}
	norm = real(cmplx.Sqrt(complex(norm, 0)))
	for i := range q.State {
		q.State[i] /= complex(norm, 0)
	}
}

// Measure performs a measurement on the qubit and returns the index of the collapsed state.
func (q *Qubit) Measure() int {
	probabilities := make([]float64, len(q.State))
	for i, amp := range q.State {
		probabilities[i] = cmplx.Abs(amp) * cmplx.Abs(amp)
	}

	outcome := rand.Intn(len(q.State))
	q.State = make([]complex128, len(q.State))
	q.State[outcome] = 1.0
	return outcome
}

// String returns a formatted representation of the qubit.
func (q *Qubit) String() string {
	return fmt.Sprintf("α|0⟩ + β|1⟩ = [%v, %v]", q.State[0], q.State[1])
}

// ApplyGate applies a quantum gate matrix to the qubit's state.
func (q *Qubit) ApplyGate(gate [][]complex128) {
	newState := make([]complex128, len(q.State))

	for i := range gate {
		newState[i] = 0
		for j := range gate[i] {
			newState[i] += gate[i][j] * q.State[j]
		}
	}

	q.State = newState
}
