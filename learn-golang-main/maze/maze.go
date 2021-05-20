package main

import (
	"GoDemoProj/encapsulation/queue"
	"fmt"
	"os"
)

var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func main() {
	maze := readMazeFromFile("./maze/maze.in")
	printMaze(maze)
	bfs(maze, point{0, 0}, point{5, 4})
}

func bfs(maze *[][]int, start, end point) {
	steps := make([][]int, len(*maze))
	for i := range *maze {
		steps[i] = make([]int, len((*maze)[i]))
	}

	q := queue.Queue{}
	q.Push(start)
	step := 1
	steps[start.x][start.y] = -1 // mark start point

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q.Poll().(point)
			for _, dir := range dirs {
				next := cur.add(dir)

				if next.isValidStep(*maze) && steps[next.x][next.y] == 0 {
					if next.x == end.x && next.y == end.y {
						steps[end.x][end.y] = step
						printMaze(&steps)
						return
					}
					steps[next.x][next.y] = step
					q.Push(next)
				}
			}
		}
		step++
	}

	return
}

func (p *point) isValidStep(grid [][]int) bool {
	if p.x < 0 || p.x >= len(grid) || p.y < 0 || p.y >= len(grid[0]) || grid[p.x][p.y] == 1 {
		return false
	}

	return true
}

func (p *point) add(dir point) point {
	return point{p.x + dir.x, p.y + dir.y}
}

type point struct {
	x, y int
}

func printMaze(maze *[][]int) {
	for _, row := range *maze {
		for _, val := range row {
			fmt.Printf("%3d", val) // 三位对齐
		}
		println()
	}
	println()
}

func readMazeFromFile(filename string) *[][]int {
	file, _ := os.Open(filename) // todo: error handling
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
	return &maze
}
