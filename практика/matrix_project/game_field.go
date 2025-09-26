// game_field.go
package main

import (
	"fmt"
	"math/rand"

	"matrix_project/matrixops"
)

func generateTerrain(size int) matrixops.Matrix {
	terrain := matrixops.New(size, size)
	for i := range terrain {
		for j := range terrain[i] {
			terrain[i][j] = rand.Float64()
		}
	}
	return terrain
}

func applyHeightmap(terrain matrixops.Matrix) {
	heightmap := matrixops.New(len(terrain), len(terrain[0]))
	for i := range heightmap {
		for j := range heightmap[i] {
			heightmap[i][j] = 0.3*terrain[i][j] + 0.7*terrain[j][i]
		}
	}

	fmt.Println("Карта высот:")
	for _, row := range heightmap {
		for _, val := range row {
			fmt.Printf("%.2f ", val)
		}
		fmt.Println()
	}
}

func main() {
	terrain := generateTerrain(4)
	applyHeightmap(terrain)
}
