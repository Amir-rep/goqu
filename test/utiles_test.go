package test

import (
	"goqu/internal/utils"
	"testing"
)

func TestMatrixCreation(t *testing.T) {
	data := []complex128{
		1 + 0i, 0 + 0i,
		0 + 0i, 1 + 0i,
	}

	matrix := utils.NewMatrix(2, 2, data)

	if matrix.Rows() != 2 {
		t.Errorf("Expected 2 rows, got %d", matrix.Rows())
	}
	if matrix.Columns() != 2 {
		t.Errorf("Expected 2 columns, got %d", matrix.Columns())
	}
	if matrix.At(0, 0) != 1+0i {
		t.Errorf("Expected matrix[0,0] to be 1+0i, got %v", matrix.At(0, 0))
	}
	if matrix.At(1, 1) != 1+0i {
		t.Errorf("Expected matrix[1,1] to be 1+0i, got %v", matrix.At(1, 1))
	}
}

//testing
