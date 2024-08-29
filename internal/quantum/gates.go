package goqu

import (
	"goqu/internal/utils"
	"math"
)

// Gate represents a quantum gate, implemented as a matrix.
type Gate struct {
	matrix *utils.Matrix
}

// NewGate creates a new quantum gate from a given matrix.
func NewGate(matrix *utils.Matrix) *Gate {
	return &Gate{matrix: matrix}
}

// Apply applies the quantum gate to a single-qubit or multi-qubit state.
// For multi-qubit states, the qubitIndex parameter specifies which qubit the gate acts on.
func (g *Gate) Apply(s *State, qubitIndex int) {
	if s.numQubits == 1 {
		// Single-qubit state, just apply the gate directly
		newState := g.matrix.Multiply(utils.NewMatrix(len(s.vector), 1, s.vector))
		s.vector = newState.Data() // Use Data() method to access the matrix data
	} else {
		// Apply the gate to a specific qubit in a multi-qubit system
		ApplySingleQubitGate(s, g.matrix, qubitIndex)
	}
}

// ApplySingleQubitGate applies a single-qubit gate to a specific qubit in a multi-qubit state.
func ApplySingleQubitGate(s *State, gate *utils.Matrix, qubitIndex int) {
	totalQubits := s.numQubits
	dim := 1 << totalQubits // 2^totalQubits

	identity := utils.IdentityMatrix(2)
	fullGate := gate

	for i := 0; i < totalQubits; i++ {
		if i == qubitIndex {
			continue
		}
		fullGate = utils.TensorProduct(identity, fullGate)
	}

	newStateVector := make([]complex128, dim)
	for i := 0; i < dim; i++ {
		amplitude := complex128(0)
		for j := 0; j < dim; j++ {
			amplitude += fullGate.At(i, j) * s.vector[j]
		}
		newStateVector[i] = amplitude
	}

	s.vector = newStateVector
}

// ApplyMultiQubitGate applies a multi-qubit gate to the entire quantum state.
// This assumes the gate matrix has the correct dimensions (e.g., 4x4 for 2 qubits).
func ApplyMultiQubitGate(s *State, gate *utils.Matrix) {
	if gate.Rows() != len(s.vector) || gate.Columns() != len(s.vector) {
		panic("Gate dimensions do not match the state vector dimensions")
	}

	newState := gate.Multiply(utils.NewMatrix(len(s.vector), 1, s.vector))
	s.vector = newState.Data() // Use Data() method to access the matrix data
}

// Standard Quantum Gates
var (
	PauliX = NewGate(utils.NewMatrix(2, 2, []complex128{
		0, 1,
		1, 0,
	}))

	PauliY = NewGate(utils.NewMatrix(2, 2, []complex128{
		0, -1i,
		1i, 0,
	}))

	PauliZ = NewGate(utils.NewMatrix(2, 2, []complex128{
		1, 0,
		0, -1,
	}))

	Hadamard = NewGate(utils.NewMatrix(2, 2, []complex128{
		1 / complex(math.Sqrt(2), 0), 1 / complex(math.Sqrt(2), 0),
		1 / complex(math.Sqrt(2), 0), -1 / complex(math.Sqrt(2), 0),
	}))

	// Identity Gate is equivalent to no operation on the state.
	Identity = NewGate(utils.IdentityMatrix(2))

	// CNOT Gate (Controlled-NOT), 2-qubit gate
	CNOT = NewGate(utils.NewMatrix(4, 4, []complex128{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 0, 1,
		0, 0, 1, 0,
	}))

	// Custom or more gates can be added here (e.g., T gate, S gate, SWAP, etc.)
)
