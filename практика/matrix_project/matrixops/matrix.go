// matrixops/matrix.go
package matrixops

import "errors"

type Matrix [][]float64

func New(rows, cols int) Matrix {
	m := make(Matrix, rows)
	for i := range m {
		m[i] = make([]float64, cols)
	}
	return m
}

func (m Matrix) Add(n Matrix) (Matrix, error) {
	if len(m) != len(n) || len(m[0]) != len(n[0]) {
		return nil, errors.New("matrix dimensions mismatch")
	}
	result := New(len(m), len(m[0]))
	for i := range m {
		for j := range m[i] {
			result[i][j] = m[i][j] + n[i][j]
		}
	}
	return result, nil
}

func (m Matrix) Multiply(n Matrix) (Matrix, error) {
	if len(m[0]) != len(n) {
		return nil, errors.New("incompatible matrix dimensions")
	}
	result := New(len(m), len(n[0]))
	for i := range m {
		for j := range n[0] {
			for k := range n {
				result[i][j] += m[i][k] * n[k][j]
			}
		}
	}
	return result, nil
}

func (m Matrix) Transpose() Matrix {
	result := New(len(m[0]), len(m))
	for i := range m {
		for j := range m[i] {
			result[j][i] = m[i][j]
		}
	}
	return result
}
