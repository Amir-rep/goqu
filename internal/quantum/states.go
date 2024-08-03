package quantum

//imports
import (
	"math"
	"math/cmplx"
)


type State struct {
	alpha complex128 //the coefficient of the zero state
	beta  complex128 // the coefficient of the one state
}

func Set_qubit(alpha complex128, beta complex128) *State{
	if math.Pow(cmplx.Abs(alpha),2) + math.Pow(cmplx.Abs(beta),2) != 1{
		panic("The qubit state must be normalized such that |alpha|^2 + |beta|^2 = 1")
	}
	return &State{Alpha: alpha, Beta: beta}
}
//next function
