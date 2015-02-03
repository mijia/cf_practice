package re

type RePattern struct {
	pattern []byte
	graph   *Digraph
}

func (re *RePattern) Match(text string) bool {
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

func (re *RePattern) buildGraph() {
	re.graph = NewDigraph(len(re.pattern) + 1)
	opStack := NewStack(len(re.pattern))
	for i := 0; i < len(re.pattern); i++ {
		leftParen := i
		curByte := re.pattern[i]
		if curByte == '(' || curByte == '|' {
			opStack.Push(i)
		} else if curByte == ')' {
			maybeOr := opStack.Pop()
			if re.pattern[maybeOr] == '|' {
				leftParen = opStack.Pop()
				re.graph.AddEdge(leftParen, maybeOr+1)
				re.graph.AddEdge(maybeOr, i)
			} else {
				leftParen = maybeOr // it is a (
			}
		}

		if i < len(re.pattern)-1 && re.pattern[i+1] == '*' {
			re.graph.AddEdge(leftParen, i+1)
			re.graph.AddEdge(i+1, leftParen)
		}
		if curByte == '(' || curByte == '*' || curByte == ')' {
			re.graph.AddEdge(i, i+1)
		}
	}
}

func NewRePattern(pattern string) *RePattern {
	data := []byte(pattern)
	re := &RePattern{
		pattern: data,
	}
	re.buildGraph()
	return re
}

type Stack struct {
	data []int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		data: make([]int, 0, capacity),
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(value int) {
	s.data = append([]int{value}, s.data...)
}

func (s *Stack) Pop() int {
	v := s.data[0]
	s.data = s.data[1:]
	return v
}

type Digraph struct {
	verticeCount int
	matrix       map[int][]int
}

func NewDigraph(verticeCount int) *Digraph {
	return &Digraph{
		verticeCount: verticeCount,
		matrix:       make(map[int][]int),
	}
}

func (dg *Digraph) AddEdge(i, j int) {
	if _, ok := dg.matrix[i]; !ok {
		dg.matrix[i] = make([]int, 0, dg.verticeCount)
	}
	dg.matrix[i] = append(dg.matrix[i], j)
}

func (dg *Digraph) Reached(starts ...int) []int {
	if len(starts) == 0 {
		return []int{}
	}
	visited := make(map[int]struct{})
	s := NewStack(dg.verticeCount)
	for _, start := range starts {
		s.Push(start)
	}
	for !s.IsEmpty() {
		node := s.Pop()
		if _, seen := visited[node]; !seen {
			visited[node] = struct{}{}
			if children, ok := dg.matrix[node]; ok {
				for _, child := range children {
					s.Push(child)
				}
			}
		}
	}
	leaves := make([]int, 0, len(visited))
	for key := range visited {
		leaves = append(leaves, key)
	}
	return leaves
}
