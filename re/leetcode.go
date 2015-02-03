package re

type LeetCodePattern struct {
	pattern []byte
	graph   *Digraph
}

func (re *LeetCodePattern) Match(text string) bool {
	data := []byte(text)
	reachable := re.graph.Reached(0)
	for i := 0; i < len(data); i++ {
		matched := make(map[int]struct{})
		for _, rState := range reachable {
			if rState == len(re.pattern) {
				continue
			}
			if re.pattern[rState] == data[i] || re.pattern[rState] == '.' {
				matched[rState+1] = struct{}{}
			}
		}
		starts := make([]int, 0, len(matched))
		for v := range matched {
			starts = append(starts, v)
		}
		reachable = re.graph.Reached(starts...)
	}
	for _, rState := range reachable {
		if rState == len(re.pattern) {
			return true
		}
	}
	return false
}

func (re *LeetCodePattern) buildGraph() {
	re.graph = NewDigraph(len(re.pattern) + 1)
	opStack := NewStack(len(re.pattern))
	for i := 0; i < len(re.pattern); i++ {
		leftParen := i
		curByte := re.pattern[i]
		if curByte == '(' {
			opStack.Push(i)
		} else if curByte == ')' {
			leftParen = opStack.Pop()
		}

		if i < len(re.pattern)-1 && re.pattern[i+1] == '*' {
			re.graph.AddEdge(leftParen, i+1)
			re.graph.AddEdge(i+1, leftParen)
		}
		if curByte == '*' || curByte == '(' || curByte == ')' {
			re.graph.AddEdge(i, i+1)
		}
	}
}

func NewLeetCodePattern(pattern string) *LeetCodePattern {
	re := &LeetCodePattern{
		pattern: []byte(pattern),
	}
	re.buildGraph()
	return re
}
