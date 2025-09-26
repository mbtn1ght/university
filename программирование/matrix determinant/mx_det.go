package main

import (
	"fmt"
	"math/rand"
)

func main() {
	matrix := makeMatrix()
	Flag := ifDet(matrix)
	if Flag == true {
		fmt.Printf("%d", det(matrix))
	}

}
func makeMatrix() [][]int {

	var column int
	var raw int
	fmt.Printf("input number of matrix's columns ")
	fmt.Scan(&column)
	matrix := make([][]int, column)
	fmt.Printf("input number of matrix's raws: ")
	fmt.Scan(&raw)

	for i := range matrix {
		matrix[i] = make([]int, raw)
	}

	for i := 0; i < column; i++ {
		for j := 0; j < raw; j++ {
			num := rand.Intn(25)
			matrix[i][j] = num
		}
	}
	for line := 0; line < column; line++ {
		fmt.Println(matrix[line])
	}
	return matrix
}
func ifDet(matrix [][]int) bool {
	var Flag bool = true
	if len(matrix) != len(matrix[0]) {
		fmt.Print("This matrix have no determinant")
		Flag = false
	}
	return Flag

}

func det(a [][]int) int {
	if len(a) == 1 {
		return a[0][0]
	}
	sign, d := 1, int(0)
	for i, x := range a[0] {
		d += int(sign) * x * det(excludeColumn(a[1:], i))
		sign *= -1
	}
	return d
}

func excludeColumn(a [][]int, i int) [][]int {
	b := make([][]int, len(a))
	n := len(a[0]) - 1
	for j, row := range a {
		r := make([]int, n)
		copy(r[:i], row[:i])
		copy(r[i:], row[i+1:])
		b[j] = r
	}
	return b
}
