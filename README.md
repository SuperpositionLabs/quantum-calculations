# Quantum Calculations

![GitHub](https://img.shields.io/github/license/SuperpositionLabs/quantum-calculations)
![GitHub last commit](https://img.shields.io/github/last-commit/SuperpositionLabs/quantum-calculations)
![GitHub repo size](https://img.shields.io/github/repo-size/SuperpositionLabs/quantum-calculations)
[![Docs](https://img.shields.io/badge/Go-Docs-blue)](https://pkg.go.dev/github.com/SuperpositionLabs/quantum-calculations)

**Quantum Calculations** is a Go library dedicated to fundamental operations in quantum computing. This project aims to provide basic implementations of qubits, quantum gates, and associated operations, serving as a foundation for further studies and developments in the field of quantum computing.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Qubit Initialization](#qubit-initialization)
  - [Applying Quantum Gates](#applying-quantum-gates)
  - [Measuring Qubits](#measuring-qubits)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Qubit Representation**: Implementation of qubits with complex states and methods for normalization and measurement.
- **Quantum Gates**: Implementation of various standard quantum gates, including:
  - **Pauli-X (NOT) Gate**: Inverts the state of a qubit.
  - **Pauli-Y Gate**: Applies a specific rotation to the qubit's state.
  - **Pauli-Z Gate**: Applies a rotation around the Z-axis to the qubit's state.
  - **Hadamard Gate**: Places the qubit into an equal superposition of the |0⟩ and |1⟩ states.
  - **Controlled-NOT (CNOT) Gate**: Applies the NOT operation to the target qubit if the control qubit is in the |1⟩ state.
  - **S Gate**: Applies a phase of π/2 to the |1⟩ state of the qubit.
  - **T Gate**: Applies a phase of π/4 to the |1⟩ state of the qubit.
- **Quantum Operations**: Functions to apply quantum operations on qubits, such as the application of the CNOT gate to control and target qubits.

## Installation

To integrate the Quantum Calculations library into your Go project, run:

```bash
go get github.com/SuperpositionLabs/quantum-calculations/pkg/qubits
go get github.com/SuperpositionLabs/quantum-calculations/pkg/gates
go get github.com/SuperpositionLabs/quantum-calculations/pkg/operations
```


Ensure that your project is using Go modules by initializing it with `go mod init`.

## Usage

### Qubit Initialization

Create and normalize a new qubit:

```go
package main

import (
    "fmt"
    "quantum-calculations/pkg/qubits"
)

func main() {
    state := []complex128{1, 1}
    qubit := qubits.NewQubit(state)
    fmt.Println(qubit)
}
```

### Applying Quantum Gates

Apply a quantum gate to a qubit:

```go
package main

import (
    "fmt"
    "quantum-calculations/pkg/qubits"
    "quantum-calculations/pkg/gates"
)

func main() {
    state := []complex128{1, 0}
    qubit := qubits.NewQubit(state)
    qubit.ApplyGate(gates.Hadamard())
    fmt.Println(qubit)
}
```

### Measuring Qubits

Measure the state of a qubit:

```go
package main

import (
    "fmt"
    "quantum-calculations/pkg/qubits"
)

func main() {
    state := []complex128{1, 0}
    qubit := qubits.NewQubit(state)
    result := qubit.Measure()
    fmt.Printf("Measurement result: %d\n", result)
}
```

## Contributing

Contributions are welcome! To contribute to the project:

1. Fork this repository.
2. Create a branch for your feature or fix (`git checkout -b feature/new-feature`).
3. Commit your changes (`git commit -m 'Add new feature'`).
4. Push to the branch (`git push origin feature/new-feature`).
5. Open a Pull Request.

Please ensure that your code adheres to the project's coding standards and includes appropriate tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

This `README.md` will be updated as the project progresses and new features are implemented. Feel free to explore, use, and contribute to the development of this quantum computing library in Go.
