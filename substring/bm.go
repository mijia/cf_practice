package substring

type BmPattern struct {
	pattern   string
	skipTable [256]int
}

func (bm BmPattern) Match(text string) int {
	data := []byte(text)
	pattern := []byte(bm.pattern)
	skip := 0
	for i := 0; i < len(data)-len(pattern); i += skip {
		skip = 0
		for j := len(pattern) - 1; j >= 0; j -= 1 {
			if pattern[j] != data[i+j] {
				skip = j - bm.skipTable[data[i+j]]
				if skip < 1 {
					skip = 1
				}
				break
			}
		}
		if skip == 0 {
			return i
		}
	}
	return len(data)
}

func NewBmPattern(pattern string) BmPattern {
	data := []byte(pattern)
	bm := BmPattern{
		pattern: pattern,
	}
	for i := range bm.skipTable {
		bm.skipTable[i] = -1
	}
	for j, x := range data {
		bm.skipTable[x] = j
	}
	return bm
}
