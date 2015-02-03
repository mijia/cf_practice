package re

import "testing"

type Case struct {
	text    string
	pattern string
	result  bool
}

func TestLeetCodePattern(t *testing.T) {
	cases := []Case{
		Case{"aa", "a", false},
		Case{"aa", "aa", true},
		Case{"aaa", "aa", false},
		Case{"aa", "a*", true},
		Case{"aa", ".*", true},
		Case{"ab", ".*", true},
		Case{"aab", "c*a*b", true},
	}
	for _, c := range cases {
		re := NewLeetCodePattern("(" + c.pattern + ")")
		matched := re.Match(c.text)
		if matched != c.result {
			t.Errorf("Leet Code Match wrong: returned %v, should be %v", matched, c.result)
		}
	}
}

func TestRePattern(t *testing.T) {
	re := NewRePattern("((A*B|AC)D)")
	cases := map[string]bool{
		"AABD": true,
		"AABC": false,
	}
	for key, value := range cases {
		matched := re.Match(key)
		if matched != value {
			t.Errorf("Match result wrong: returned %v, should be %v", matched, value)
		}
	}
}

func TestStack(t *testing.T) {
	s := NewStack(10)
	for i := 1; i <= 5; i++ {
		s.Push(i)
	}
	for i := 5; i >= 1; i-- {
		d := s.Pop()
		if d != i {
			t.Errorf("Stack push pop not working, %d != %d", i, d)
		}
	}
}

func TestDirectDFS(t *testing.T) {
	graph := NewDigraph(4)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(2, 3)

	visited := graph.Reached(0)
	if len(visited) != 4 {
		t.Errorf("Direct graph dfs failed, returned %v", visited)
	}
}
