package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

func (p point) add(another point) point {
	return point{p.i + another.i, p.j + another.j}
}

func (p point) from(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) || p.j < 0 || p.j >= len(grid[0]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func main() {
	maze := getMaze("maze/maze.in")
	path := getPath(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range path {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func getPath(maze [][]int, start point, end point) (path [][]int) {
	path = make([][]int, len(maze))
	for i := range maze {
		path[i] = make([]int, len(maze[i]))
	}

	toExplore := []point{start}
	for len(toExplore) > 0 {
		cur := toExplore[0]
		toExplore = toExplore[1:]

		for _, dir := range dirs {
			next := cur.add(dir)

			val, ok := next.from(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.from(path)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			step, _ := cur.from(path)
			path[next.i][next.j] = step + 1
			if next == end {
				return
			}

			toExplore = append(toExplore, next)
		}
	}
	return
}

func getMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}
