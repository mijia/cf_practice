package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var _ = fmt.Println

type IntSorter struct {
	data []int
}

func (s *IntSorter) Len() int           { return len(s.data) }
func (s *IntSorter) Swap(i, j int)      { s.data[i], s.data[j] = s.data[j], s.data[i] }
func (s *IntSorter) Less(i, j int) bool { return s.data[i] < s.data[j] }

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	N := ReadInt32()
	boys := make([]int, N)
	for i := 0; i < N; i += 1 {
		boys[i] = ReadInt32()
	}

	M := ReadInt32()
	girls := make([]int, M)
	for i := 0; i < M; i += 1 {
		girls[i] = ReadInt32()
	}

	bSorter := &IntSorter{boys}
	gSorter := &IntSorter{girls}
	sort.Sort(bSorter)
	sort.Sort(gSorter)
	boys, girls = bSorter.data, gSorter.data

	bi, gi := 0, 0
	matches := 0
	for bi < N && gi < M {
		if boys[bi] == girls[gi] || boys[bi]-girls[gi] == -1 || boys[bi]-girls[gi] == 1 {
			matches++
			bi++
			gi++
			continue
		}
		if boys[bi] > girls[gi] {
			gi++
		} else {
			bi++
		}
	}
	PrintInts(matches)
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
