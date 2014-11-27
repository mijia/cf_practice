package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var _ = fmt.Println

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	forwardPos := make(map[int]int)
	backwardPos := make(map[int]int)

	N := ReadInt32()
	position := make([]int, N)
	for i := 0; i < N; i += 1 {
		x, y := ReadInt32(), ReadInt32()
		forwardPos[x] = y
		backwardPos[y] = x
	}

	startId := 0
	startPos := 1
	for {
		if backId, ok := forwardPos[startId]; ok && startPos < N {
			position[startPos] = backId
			startId = backId
			startPos += 2
		} else {
			break
		}
	}
	for fId := range forwardPos {
		if _, ok := backwardPos[fId]; !ok {
			startId = fId
			break
		}
	}
	position[0] = startId
	startPos = 2
	for {
		if backId, ok := forwardPos[startId]; ok && startPos < N {
			position[startPos] = backId
			startId = backId
			startPos += 2
		} else {
			break
		}
	}
	PrintInts(position...)
}

func ReadInt32() int {
	scanner.Scan()
	ans, _ := strconv.Atoi(scanner.Text())
	return ans
}

func PrintInts(ints ...int) {
	for _, value := range ints {
		writer.WriteString(strconv.Itoa(value))
		writer.WriteByte(' ')
	}
}
