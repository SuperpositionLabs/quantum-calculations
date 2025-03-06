package qubits

import (
	"fmt"
	"math/cmplx"
	"math/rand"
)

type Qubit struct {
	State []complex128
}

// Returns an normalized qubit
func NewQubit(state []complex128) *Qubit {
	q := &Qubit{State: state}
	q.Normalize()
	return q
}

// Returns an non normalized qubit
func NewNNQubit(state []complex128) *Qubit {
	q := &Qubit{State: state}
	return q
}

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

func (q *Qubit) String() string {
	return fmt.Sprintf("α|0⟩ + β|1⟩ = [%v, %v]", q.State[0], q.State[1])
}

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
