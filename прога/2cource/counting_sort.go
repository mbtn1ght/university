package main

import "fmt"

func main() {
	array := []int{23, 4, 49, 8, 12, 5}
	n := len(array)
	count := make([]int, n)
	res := make([]int, n)

	counted := make([][]bool, n)
	for i := 0; i < n; i++ {
		counted[i] = make([]bool, n)
	}

	fmt.Println("Исходный массив:", array)
	fmt.Println("Начальные значения счетчиков:", count)

	step := 1
	for {
		any := false
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if array[j] < array[i] && !counted[i][j] {
					counted[i][j] = true
					count[i]++
					any = true
					break
				}
			}
		}
		if !any {
			break
		}
		fmt.Printf("Значения счетчиков после %d-го цикла: %v\n", step, count)
		step++
	}

	fmt.Printf("Итоговые счётчики: %v\n", count)

	occupied := make([]bool, n)
	for i := 0; i < n; i++ {
		pos := count[i]
		for pos < n && occupied[pos] {
			pos++
		}
		if pos >= n {
			for k := 0; k < n; k++ {
				if !occupied[k] {
					pos = k
					break
				}
			}
		}
		res[pos] = array[i]
		occupied[pos] = true
	}

	fmt.Println("Отсортированный массив:", res)
}
