package main

import "fmt"

func makeMatrix(n int, m int) [][]int {
	matrix := make([][]int, n)

	for i := range matrix {
		matrix[i] = make([]int, m)
	}
	fmt.Println("Enter elements : ")
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Scanf("%d", &matrix[i][j])
		}
	}
	fmt.Println(matrix)
	return matrix
}

func main() {
	var (
		n, m int
	)
	fmt.Println("Number of lines : ")
	fmt.Scan(&n)
	fmt.Println("Number of columns: ")
	fmt.Scan(&m)
	makeMatrix(n, m)

}
