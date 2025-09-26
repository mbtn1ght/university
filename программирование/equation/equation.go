package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	var a, b, c float64
	fmt.Print("input the coefficient of the equation.")
	fmt.Println("input the coefficient 'a': ")
	fmt.Scan(&a)
	fmt.Println("input the coefficient 'b': ")
	fmt.Scan(&b)
	fmt.Println("input the coefficient 'c': ")
	fmt.Scan(&c)
	findSolution(a, b, c)

}
func findSolution(a, b, c float64) {
	var D float64
	D = b*b - 4*a*c
	if D > 0 {
		var x1, x2 float64
		x1 = ((b * (-1)) + math.Sqrt(D)) / (2 * a)
		x2 = ((b * (-1)) - math.Sqrt(D)) / (2 * a)
		fmt.Print("Quadratic equation has two roots: ", x1, " and ", x2)
	} else if D == 0 {
		var x float64
		x = (b * (-1)) / (2 * a)
		fmt.Print("Quadratic equation has one root: ", x)
	} else {
		var FirstPart, SecondPart float64
		FirstPart = (b * (-1)) / (2 * a)
		SecondPart = math.Sqrt(math.Abs(D)) / (2 * a)
		x1 := complex(FirstPart, SecondPart)
		x2 := complex(FirstPart, SecondPart*(-1))
		fmt.Println("Quadratic equation has two roots: ")
		fmt.Println(cmplx.Conj(x1))
		fmt.Println(cmplx.Conj(x2))
	}

}
