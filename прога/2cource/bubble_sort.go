package main

import "fmt"

func main() {
	slice := []int{9, 4, 3, 2, 6, 1, 5, 8, 7, 0}

	for k := 0; ; k++ {
		count := 0
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				count++
			}
		}

		fmt.Printf("%v проход: ", k)
		for _, v := range slice {
			fmt.Printf("%d ", v)
		}
		fmt.Println()

		if count == 0 {
			break
		}
	}

	fmt.Print("Результат: ")
	for _, v := range slice {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
