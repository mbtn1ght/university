package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {

	file, err := os.Open("labyrinth.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	var maze [][]rune
	scanner := bufio.NewScanner(file)
	start := Point{-1, -1}
	end := Point{-1, -1}

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		row := []rune(line)
		maze = append(maze, row)
		for x, char := range row {
			if char == 'S' {
				start = Point{x, y}
			} else if char == 'E' {
				end = Point{x, y}
			}
		}
	}

	if start.x == -1 || start.y == -1 {
		fmt.Println("Не найдена начальная точка 'S'")
		return
	}
	if end.x == -1 || end.y == -1 {
		fmt.Println("Не найдена конечная точка 'E'")
		return
	}

	path, found := bfs(maze, start, end)
	if !found {
		fmt.Println("Путь от 'S' до 'E' не найден")
		return
	}

	for _, p := range path {
		if maze[p.y][p.x] != 'S' && maze[p.y][p.x] != 'E' {
			maze[p.y][p.x] = '*'
		}
	}

	for _, row := range maze {
		fmt.Println(string(row))
	}

	fmt.Println("Длина пройденного пути:", len(path)-1)
}

func bfs(maze [][]rune, start, end Point) ([]Point, bool) {
	queue := [][]Point{{start}}
	visited := make(map[Point]bool)
	visited[start] = true

	directions := []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		current := path[len(path)-1]

		if current == end {
			return path, true
		}

		for _, d := range directions {
			next := Point{current.x + d.x, current.y + d.y}

			if next.y < 0 || next.y >= len(maze) || next.x < 0 || next.x >= len(maze[next.y]) {
				continue
			}

			if maze[next.y][next.x] == '#' || visited[next] {
				continue
			}

			visited[next] = true
			newPath := make([]Point, len(path))
			copy(newPath, path)
			newPath = append(newPath, next)
			queue = append(queue, newPath)
		}
	}

	return nil, false
}
