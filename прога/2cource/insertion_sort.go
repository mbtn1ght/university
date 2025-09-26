package main

import "fmt"

func main() {
	slice := []int{5, 2, 4, 1, 3}

	fmt.Println("Изначальный массив:", slice)

	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1

		// ищем позицию для вставки
		for j >= 0 && slice[j] > key {
			j--
		}

		// сдвигаем блок вправо
		copy(slice[j+2:i+1], slice[j+1:i])

		// вставляем key
		slice[j+1] = key

		// массив после цикла
		fmt.Printf("Массив после %d цикла: %v\n", i, slice)
	}

	fmt.Println("Отсортированный массив:", slice)
}
