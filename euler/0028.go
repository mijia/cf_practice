package main

import "fmt"

func main() {
	n := 1001
	matrix := initSpiralMatrix(n)
	sum := 0
	for i := 0; i < n; i += 1 {
		sum += matrix[i][i]
		sum += matrix[i][n-1-i]
	}
	sum -= 1
	fmt.Println(sum)
}

const (
	goDown = iota
	goLeft
	goUp
	goRight
)

func initSpiralMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := 0; i < n; i += 1 {
		m[i] = make([]int, n)
	}
	curDigit := 1
	cx, cy := n/2, n/2
	m[cy][cx] = curDigit
	curDigit += 1
	for circle := 3; circle <= n; circle += 2 {
		stepCounts := circle*circle - (circle-2)*(circle-2)
		dirs := make([]int, 0, stepCounts)
		dirs = append(dirs, goRight)
		for i := 0; i < 4; i += 1 {
			addCounts := circle - 1
			if i == 0 {
				addCounts = circle - 2
			}
			for j := 0; j < addCounts; j += 1 {
				dirs = append(dirs, i)
			}
		}
		for step := 0; step < stepCounts; step += 1 {
			switch dirs[step] {
			case goDown:
				cy += 1
			case goLeft:
				cx -= 1
			case goUp:
				cy -= 1
			case goRight:
				cx += 1
			}
			m[cy][cx] = curDigit
			curDigit += 1
		}
	}
	return m
}
