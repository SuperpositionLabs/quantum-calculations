package main

import (
	"fmt"

	"quantum-calculations/pkg/qubits"
)

func main() {
	state := []complex128{1, 1} // Estado |0> + |1>
	qubit := qubits.NewQubit(state)
	fmt.Println("Estado inicial do qubit: ", qubit)

	outcome := qubit.Measure()
	fmt.Println("Resultado da medição", outcome)
	fmt.Println("Estado do qubit após medição", qubit)
}
