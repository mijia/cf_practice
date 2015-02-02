package substring

import "testing"

func TestDfaMatching(t *testing.T) {
	dfaPattern := NewDfaPattern("ABABAC")
	if matchedPos := dfaPattern.Match("AABACAABABACAA"); matchedPos != 6 {
		t.Errorf("Should find string at index 6 but returned %d", matchedPos)
	}
}

func TestKmpMatching(t *testing.T) {
	kmpPattern := NewKmpPattern("ABABAC")
	if matchedPos := kmpPattern.Match("AABACAABABACAA"); matchedPos != 6 {
		t.Errorf("Should find string at index 6 but returned %d", matchedPos)
	}
}
