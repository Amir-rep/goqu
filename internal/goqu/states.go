package goqu

//imports
import (
	"goqu/internal/utils"
	"math"
	"math/cmplx"
)

type State struct {
	alpha complex128 //the coefficient of the zero state
	beta  complex128 // the coefficient of the one state
}

func Set_qubit(alpha complex128, beta complex128) *State {
	if math.Pow(cmplx.Abs(alpha), 2)+math.Pow(cmplx.Abs(beta), 2) != 1 {
		panic("The qubit state must be normalized such that |alpha|^2 + |beta|^2 = 1")
	}
	return &State{alpha: alpha, beta: beta}
}

func Get_one_prob(s State) float64 {
	return math.Pow(cmplx.Abs(s.alpha), 2)
}

func Get_zero_prob(s State) float64 {
	return math.Pow(cmplx.Abs(s.beta), 2)
}

//turn the state into vector form

func (s *State) Vector_state() *utils.Matrix {
	data := []complex128{s.alpha, s.beta}
	return utils.New_Matrix(2, 1, data)
}

//next function
