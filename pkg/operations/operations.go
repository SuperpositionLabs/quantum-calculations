package operations

import (
	"github.com/SuperpositionLabs/quantum-calculations/pkg/gates"
	"github.com/SuperpositionLabs/quantum-calculations/pkg/qubits"
)

// ApplyCNOT applies a Controlled-NOT (CNOT) gate to a pair of qubits.
// The control qubit determines whether the target qubit is flipped.
func ApplyCNOT(control *qubits.Qubit, target *qubits.Qubit) {
	// Combine the states of the control and target qubits into a 4-element state vector.
	combinedState := make([]complex128, 4)
	combinedState[0] = control.State[0] * target.State[0]
	combinedState[1] = control.State[0] * target.State[1]
	combinedState[2] = control.State[1] * target.State[0]
	combinedState[3] = control.State[1] * target.State[1]

	// Retrieve the CNOT gate matrix.
	gate := gates.CNOT()
	newState := make([]complex128, 4)

	// Apply the CNOT gate to the combined state.
	for i := range gate {
		newState[i] = 0
		for j := range gate[i] {
			newState[i] += gate[i][j] * combinedState[j]
		}
	}

	// Update the states of the control and target qubits based on the transformed state.
	control.State[0] = newState[0] + newState[1]
	control.State[1] = newState[2] + newState[3]
	target.State[0] = newState[0] + newState[2]
	target.State[1] = newState[1] + newState[3]

	// Normalize the qubits to maintain valid quantum states.
	control.Normalize()
	target.Normalize()
}
