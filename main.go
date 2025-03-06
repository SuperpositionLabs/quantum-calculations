package main

import (
	"fmt"
	"math"
	"math/cmplx"

	"quantum-calculations/pkg/gates"
	"quantum-calculations/pkg/operations"
	"quantum-calculations/pkg/qubits"
)

func main() {
	state := []complex128{1, 1} // Estado |0> + |1>
	qubit := qubits.NewQubit(state)
	fmt.Println("Estado inicial do qubit: ", qubit)

	outcome := qubit.Measure()
	fmt.Println("Resultado da medição", outcome)
	fmt.Println("Estado do qubit após medição", qubit)
	// -------------- PORTA CNOT ----------------------
	controlState := []complex128{1, 0}
	targetState := []complex128{0, 1}
	control := qubits.NewQubit(controlState)
	target := qubits.NewQubit(targetState)
	fmt.Println("-----------------Porta CNOT -------------")
	fmt.Println("Estado inicial do qubit de controle: ", control)
	fmt.Println("Estado inicial do qubit alvo: ", target)
	// CNOT TESTE

	operations.ApplyCNOT(control, target)
	fmt.Println("Estado após aplicar a porta CNOT:")
	fmt.Println("Qubit de controle: ", control)
	fmt.Println("Qubit alvo: ", target)

	// -------------- Porta T ----------------------
	fmt.Println("-----------------Porta T ----------------")
	runTGateTest()
}

func runTGateTest() {
	state := []complex128{0, 1}
	qubit := qubits.NewQubit(state)

	qubit.ApplyGate(gates.T())

	fmt.Println("Estado após aplicar a porta T:", qubit)

	expected := []complex128{0, cmplx.Exp(1i * (math.Pi / 4))}
	for i := range qubit.State {
		if qubit.State[i] != expected[i] {
			fmt.Printf("Erro: \n Esperado: %v \n Obtido: %v", expected, qubit.State)
			return
		}
	}
	fmt.Println("Teste da porta T funcionou!")
}
