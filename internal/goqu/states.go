package goqu

import (
	"fmt"
	//"goqu/internal/utils"
	"math"
	"math/cmplx"
)

// State represents a multi-qubit quantum state.
type State struct {
	vector    []complex128 // The vector representing the quantum state
	numQubits int          // Number of qubits in the system
}

// NewMultiQubitState creates a new quantum state for `n` qubits, initialized to the |0...0⟩ state.
func NewMultiQubitState(numQubits int) *State {
	size := int(math.Pow(2, float64(numQubits)))
	vector := make([]complex128, size)
	vector[0] = 1 // Initialize to |0...0⟩
	return &State{
		vector:    vector,
		numQubits: numQubits,
	}
}

// NewCustomState creates a new custom multi-qubit quantum state given a state vector.
func NewCustomState(vector []complex128) *State {
	numQubits := int(math.Log2(float64(len(vector))))
	if len(vector) != int(math.Pow(2, float64(numQubits))) {
		panic("State vector length must be a power of 2")
	}
	return &State{
		vector:    vector,
		numQubits: numQubits,
	}
}

// Vector returns the vector representation of the quantum state.
func (s *State) Vector() []complex128 {
	return s.vector
}

// PrintState prints the quantum state vector.
func (s *State) PrintState() {
	for _, amplitude := range s.vector {
		fmt.Printf("%v ", amplitude)
	}
	fmt.Println()
}

// ProbabilityOf returns the probability of measuring the system in a specific state |b⟩.
func (s *State) ProbabilityOf(b int) float64 {
	if b < 0 || b >= len(s.vector) {
		panic("Invalid basis state index")
	}
	return math.Pow(cmplx.Abs(s.vector[b]), 2)
}

//following functions

