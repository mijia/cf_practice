package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	_, _ = ReadInt32(), ReadInt32()
	C := ReadInt32()
	hotelX := make([]int, C)
	hotelY := make([]int, C)
	for i := 0; i < C; i += 1 {
		hotelX[i] = ReadInt32()
		hotelY[i] = ReadInt32()
	}

	H := ReadInt32()
	restX := make([]int, H)
	restY := make([]int, H)
	for i := 0; i < H; i += 1 {
		restX[i] = ReadInt32()
		restY[i] = ReadInt32()
	}

	centerX, centerY := 0.0, 0.0
	for i := 0; i < C; i += 1 {
		centerX += float64(hotelX[i])
		centerY += float64(hotelY[i])
	}
	centerX /= float64(C)
	centerY /= float64(C)

	minDist := 9999999999.0
	minIndex := -1
	for i := 0; i < H; i += 1 {
		dist := math.Abs(float64(restX[i])-centerX) + math.Abs(float64(restY[i])-centerY)
		if dist < minDist {
			minDist = dist
			minIndex = i
		}
	}

	maxDist := 0
	for i := 0; i < C; i += 1 {
		dist := AbsInt(hotelX[i]-restX[minIndex]) + AbsInt(hotelY[i]-restY[minIndex])
		if dist >= maxDist {
			maxDist = dist
		}
	}
	PrintInts(maxDist)
	writer.WriteByte('\n')
	PrintInts(minIndex + 1)

}

func AbsInt(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
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
