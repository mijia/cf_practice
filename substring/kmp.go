package substring

type KmpPattern struct {
	pattern string
	pTable  []int
}

func (kmp KmpPattern) Match(text string) int {
	data := []byte(text)
	pattern := []byte(kmp.pattern)
	for i, j := 0, 0; i < len(data); i += 1 {
		for j >= 0 && data[i] != pattern[j] {
			j = kmp.pTable[j]
		}
		j += 1

		if j == len(pattern) {
			return i + 1 - len(pattern)
		}
	}
	return len(data)
}

func NewKmpPattern(pattern string) KmpPattern {
	data := []byte(pattern)
	pTable := make([]int, len(data)+1)
	j := -1
	pTable[0] = -1
	for i := 0; i < len(data); i += 1 {
		for j >= 0 && data[i] != data[j] {
			j = pTable[j]
		}
		j += 1
		pTable[i+1] = j
	}
	return KmpPattern{
		pattern: pattern,
		pTable:  pTable,
	}
}
