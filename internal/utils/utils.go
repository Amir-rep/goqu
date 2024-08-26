package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// Matrix struct represents a 2D matrix of complex numbers
type Matrix struct {
	rows, cols int
	data       []complex128
}

// NewMatrix creates a new matrix with the specified number of rows and columns.
// The data slice should have a length equal to rows * cols.
func NewMatrix(rows, cols int, data []complex128) *Matrix {
	if len(data) != rows*cols {
		panic("data slice length does not match dimensions")
	}
	return &Matrix{
		rows: rows,
		cols: cols,
		data: data,
	}
}

// Rows returns the number of rows in the matrix.
func (m *Matrix) Rows() int {
	return m.rows
}

// Columns returns the number of columns in the matrix.
func (m *Matrix) Columns() int {
	return m.cols
}

func (m *Matrix) Data() []complex128 {
	return m.data
}

// At returns the element at the specified position in the matrix.
func (m *Matrix) At(row, col int) complex128 {
	if row < 0 || row >= m.rows || col < 0 || col >= m.cols {
		panic("index out of range")
	}
	return m.data[row*m.cols+col]
}

// Set sets the element at the specified row and column in the matrix.
func (m *Matrix) Set(row, col int, value complex128) {
	if row < 0 || row >= m.rows || col < 0 || col >= m.cols {
		panic("index out of range")
	}
	m.data[row*m.cols+col] = value
}

// Dims returns the dimensions of the matrix.
func (m *Matrix) Dims() (int, int) {
	return m.rows, m.cols
}

// PrintMatrix is a simple helper function to display the matrix.
func (m *Matrix) PrintMatrix() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Printf("%v ", m.At(i, j))
		}
		fmt.Println()
	}
}

// Multiply multiplies the matrix by another matrix and returns the result.
func (m *Matrix) Multiply(other *Matrix) *Matrix {
	if m.cols != other.rows {
		panic("matrix dimensions do not match for multiplication")
	}
	resultData := make([]complex128, m.rows*other.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < other.cols; j++ {
			var sum complex128
			for k := 0; k < m.cols; k++ {
				sum += m.At(i, k) * other.At(k, j)
			}
			resultData[i*other.cols+j] = sum
		}
	}
	return NewMatrix(m.rows, other.cols, resultData)
}

// RandomFloat64 returns a random float64 between 0 and 1.
// This can be used for probabilistic operations such as quantum measurements.
func RandomFloat64() float64 {
	rand.NewSource(time.Now().UnixNano())
	return rand.Float64()
}

// ComplexConjugate returns the complex conjugate of a given complex number.
func ComplexConjugate(c complex128) complex128 {
	return complex(real(c), -imag(c))
}

// IdentityMatrix returns an identity matrix of the specified size.
func IdentityMatrix(size int) *Matrix {
	data := make([]complex128, size*size)
	for i := 0; i < size; i++ {
		data[i*size+i] = 1
	}
	return NewMatrix(size, size, data)
}

//Tensor product for multiple state handling

func TensorProduct(a, b *Matrix) *Matrix {
	newRows := a.Rows() * b.Rows()
	newCols := a.Columns() * b.Columns()
	newData := make([]complex128, newRows*newCols)

	for i := 0; i < a.Rows(); i++ {
		for j := 0; j < a.Columns(); j++ {
			for k := 0; k < b.Rows(); k++ {
				for l := 0; l < b.Columns(); l++ {
					newData[(i*b.Rows()+k)*newCols+(j*b.Columns()+l)] = a.At(i, j) * b.At(k, l)
				}
			}
		}
	}

	return NewMatrix(newRows, newCols, newData)
}
