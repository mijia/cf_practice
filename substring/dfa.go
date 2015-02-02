package substring

type DfaPattern struct {
	pattern string
	machine map[byte][]int
}

func (dfa DfaPattern) Match(text string) int {
	data := []byte(text)
	pattern := []byte(dfa.pattern)
	i, j := 0, 0
	for ; i < len(data) && j < len(pattern); i += 1 {
		j = dfa.machine[data[i]][j]
	}
	if j == len(pattern) {
		return i - len(pattern)
	}
	return -1
}

func NewDfaPattern(pattern string) DfaPattern {
	data := []byte(pattern)
	machine := make(map[byte][]int)
	for _, p := range data {
		if _, ok := machine[p]; !ok {
			machine[p] = make([]int, len(data))
		}
	}
	machine[data[0]][0] = 1
	X := 0
	for j := 1; j < len(data); j += 1 {
		for key, _ := range machine {
			machine[key][j] = machine[key][X]
		}
		machine[data[j]][j] = j + 1
		X = machine[data[j]][X]
	}
	return DfaPattern{
		pattern: pattern,
		machine: machine,
	}
}
