package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Point struct{ X, Y int }

func main() {
	rand.Seed(time.Now().UnixNano())

	var width, height int
	fmt.Print("Введите ширину лабиринта: ")
	fmt.Scan(&width)
	fmt.Print("Введите высоту лабиринта: ")
	fmt.Scan(&height)

	if width < 3 || height < 3 {
		fmt.Println("Минимальный размер лабиринта — 3x3.")
		return
	}

	lab := make([][]rune, height)
	for y := range lab {
		lab[y] = make([]rune, width)
		for x := range lab[y] {
			lab[y][x] = '#'
		}
	}

	startX := rand.Intn(width-2) + 1
	start := Point{X: startX, Y: 0}
	lab[start.Y][start.X] = 'S'

	current := Point{X: start.X, Y: start.Y + 1}
	lab[current.Y][current.X] = ' '

	visited := map[Point]bool{
		start:   true,
		current: true,
	}

	path := []Point{current}

	pathLength := rand.Intn(width) + height

	for i := 0; i < pathLength; i++ {
		directions := []Point{
			{X: 0, Y: 1}, {X: 0, Y: -1}, {X: 1, Y: 0}, {X: -1, Y: 0},
		}
		rand.Shuffle(len(directions), func(i, j int) {
			directions[i], directions[j] = directions[j], directions[i]
		})

		moved := false
		for _, d := range directions {
			next := Point{X: current.X + d.X, Y: current.Y + d.Y}
			if next.X <= 0 || next.X >= width-1 || next.Y <= 0 || next.Y >= height-1 {
				continue
			}
			if !visited[next] {
				lab[next.Y][next.X] = ' '
				visited[next] = true
				current = next
				path = append(path, current)
				moved = true
				break
			}
		}
		if !moved {
			break
		}
	}

	end := path[len(path)-1]
	lab[end.Y][end.X] = 'E'

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if lab[y][x] == '#' && rand.Float32() < 0.1 {
				lab[y][x] = ' '
			}
		}
	}

	file, err := os.Create("labyrinth.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	for _, row := range lab {
		for _, cell := range row {
			file.WriteString(string(cell))
		}
		file.WriteString("\n")
	}

	fmt.Println("Лабиринт успешно сгенерирован и сохранён в labyrinth.txt.")
}
