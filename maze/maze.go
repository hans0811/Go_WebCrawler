package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	// Fscanf change the value
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

type point struct {
	i, j int
}

// define direction
var dirs = [4]point{
	{-1, 0}, // up
	{0, -1}, // left
	{1, 0},  // down
	{0, 1},  // right

}

// Add aix
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// Check Point Value
func (p point) at(grid [][]int) (int, bool) {

	// out of boundary
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))

	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	// Start a queue
	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			// Find new Area plus directions
			next := cur.add(dir)

			// maze at next is 0
			// and steps at next is 0
			// and next != start
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)

			steps[next.i][next.j] = curSteps + 1

			Q = append(Q, next)

		}
	}

	return steps

}

func main() {
	maze := readMaze("maze/maze.in")

	//for _, row := range maze {
	//	for _, val := range row {
	//		fmt.Printf("%d", val)
	//	}
	//	fmt.Println()
	//}

	steps := walk(maze, point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

}
