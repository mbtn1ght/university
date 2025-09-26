// equation_solver.go
package main

import (
	"fmt"

	"matrix_project/matrixops"
)

func main() {
	// Система уравнений:
	// 2x + y = 8
	// x - y = 1
	A := matrixops.Matrix{
		{2, 1},
		{1, -1},
	}

	B := matrixops.Matrix{
		{8},
		{1},
	}

	// Вычисление обратной матрицы
	det := A[0][0]*A[1][1] - A[0][1]*A[1][0]
	inv := matrixops.Matrix{
		{A[1][1] / det, -A[0][1] / det},
		{-A[1][0] / det, A[0][0] / det},
	}

	// Решение: X = A⁻¹ * B
	solution, _ := inv.Multiply(B)
	fmt.Printf("Решение: x = %.2f, y = %.2f\n", solution[0][0], solution[1][0])
}
