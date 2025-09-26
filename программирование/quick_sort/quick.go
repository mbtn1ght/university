package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func readNumbers(filename string) ([]int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	parts := strings.Fields(string(data))
	numbers := make([]int, 0, len(parts))
	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			continue
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}

func parallelQuickSort(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1
	pivot := arr[len(arr)/2]

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	if right > 0 {
		wg.Add(1)
		go parallelQuickSort(arr[:right+1], wg)
	}
	if left < len(arr) {
		wg.Add(1)
		go parallelQuickSort(arr[left:], wg)
	}
}

func main() {
	numbers, err := readNumbers("numbers.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	fmt.Println("Исходные числа:", numbers)

	var wg sync.WaitGroup
	wg.Add(1)
	go parallelQuickSort(numbers, &wg)
	wg.Wait()
	var final []int
	for i := len(numbers) - 1; i >= 0; i-- {
		final = append(final, numbers[i])
	}

	fmt.Println("Отсортированные числа:", final)
}
