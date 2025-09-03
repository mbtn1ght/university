package main

import "fmt"

func main() {
	slice := []int{9, 4, 3, 2, 6, 1, 5, 8, 7, 0}

	for {
		count := 0
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				count += 1

			}
		}
		if count == 0 {
			break
		}
	}
	fmt.Println(slice)
}
