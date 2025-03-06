package operations

import (
	"quantum-calculations/pkg/gates"
	"quantum-calculations/pkg/qubits"
)

func ApplyCNOT(control *qubits.Qubit, target *qubits.Qubit) {
	combinedState := make([]complex128, 4)
	combinedState[0] = control.State[0] * target.State[0]
	combinedState[1] = control.State[0] * target.State[1]
	combinedState[2] = control.State[1] * target.State[0]
	combinedState[3] = control.State[1] * target.State[1]

	gate := gates.CNOT()
	newState := make([]complex128, 4)
	for i := range gate {
		newState[i] = 0
		for j := range gate[i] {
			newState[i] += gate[i][j] * combinedState[j]
		}
	}

	control.State[0] = newState[0] + newState[1]
	control.State[1] = newState[2] + newState[3]
	target.State[0] = newState[0] + newState[2]
	target.State[1] = newState[1] + newState[3]

	control.Normalize()
	target.Normalize()
}
