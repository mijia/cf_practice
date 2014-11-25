package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var _ = fmt.Println

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	initState := make([]int64, 5)
	initState[0] = 0
	for i := 1; i < 5; i += 1 {
		initState[i] = ReadInt64()
	}

	stack := make([][]int64, 1)
	stack[0] = initState
	for len(stack) > 0 {
		s := stack[0]
		if s[1] * s[2] == s[3] * s[4] {
			PrintInts(s[0])
			writer.WriteByte('\n')
			PrintInts(s[1], s[2])
			writer.WriteByte('\n')
			PrintInts(s[3], s[4])
			return
		}
		fmt.Println(s)
		states := genSolutions(s)
		stack = stack[1:]
		for _, ns := range states {
			if ns[1] * ns[2] == ns[3] * ns[4] {
				PrintInts(ns[0])
				writer.WriteByte('\n')
				PrintInts(ns[1], ns[2])
				writer.WriteByte('\n')
				PrintInts(ns[3], ns[4])
				return
			}
			if ns[1] % 2 != 0 && ns[1] % 3 != 0 &&
					ns[2] % 2 != 0 && ns[2] % 3 != 0 &&
					ns[3] % 2 != 0 && ns[3] % 3 != 0 &&
					ns[4] % 2 != 0 && ns[4] % 3 != 0 {
				continue
			}
			stack = append(stack, ns)
		}
	}
	PrintInts(-1)	
}

type Sorter struct {
	data [][]int64
}
func (s *Sorter) Len() int { return len(s.data) }
func (s *Sorter) Swap(i, j int) { s.data[i], s.data[j] = s.data[j], s.data[i] }
func (s *Sorter) Less(i, j int) bool {
	sizeA := s.data[i][1] * s.data[i][2] - s.data[i][3]*s.data[i][4]
	sizeB := s.data[j][1] * s.data[j][2] - s.data[j][3]*s.data[j][4]
	return sizeA * sizeA < sizeB * sizeB
}

func genSolutions(s []int64) [][]int64 {
	nextStates := make([][]int64, 0)
	pivotBlock := 1
	if s[1] * s[2] < s[3] * s[4] {
		pivotBlock = 3
	}
	common := make([]int64, 5)
	copy(common, s)
	common[0]++
	x, y := s[pivotBlock], s[pivotBlock + 1]
	if x % 2 == 0 {
		newState := make([]int64, 5)
		copy(newState, common)
		newState[pivotBlock] = x / 2
		nextStates = append(nextStates, newState)
	}
	if x % 3 == 0 {
		newState := make([]int64, 5)
		copy(newState, common)
		newState[pivotBlock] = x / 3 * 2
		nextStates = append(nextStates, newState)
	}
	if y % 2 == 0 {
		newState := make([]int64, 5)
		copy(newState, common)
		newState[pivotBlock + 1] = y / 2
		nextStates = append(nextStates, newState)
	}
	if y % 3 == 0 {
		newState := make([]int64, 5)
		copy(newState, common)
		newState[pivotBlock + 1] = y / 3 * 2
		nextStates = append(nextStates, newState)
	}
	sorter := &Sorter{nextStates}
	sort.Sort(sorter)
	return sorter.data
}

func ReadInt32() int {
	scanner.Scan()
	ans, _ := strconv.Atoi(scanner.Text())
	return ans
}

func ReadInt64() int64 {
	scanner.Scan()
	ans, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return ans
}

func PrintInts(ints ...int64) {
	for _, value := range ints {
		writer.WriteString(strconv.FormatInt(value, 10))
		writer.WriteByte(' ')
	}
}
