package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]string, 100)
	for i := range numbers {
		numbers[i] = fmt.Sprintf("%d", rand.Intn(201)-100)
	}

	data := strings.Join(numbers, " ")
	err := os.WriteFile("numbers.txt", []byte(data), 0644)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
		return
	}

	fmt.Println("Файл 'numbers.txt' успешно создан.")
}
