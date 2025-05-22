package set

import (
	"testing"
)

func TestSet(t *testing.T) {
	s := New()
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}

	s.Insert("a")
	if !s.Contains("a") {
		t.Errorf("Expected set to contain 'a'")
	}

	s.Insert("a")
	if s.Size() != 1 {
		t.Errorf("Expected size 1, got %d", s.Size())
	}
	if !s.Contains("a") {
		t.Errorf("Expected set to contain 'a'")
	}

	s.Insert("b")
	if s.Size() != 2 {
		t.Errorf("Expected size 2, got %d", s.Size())
	}
	if !s.Contains("b") {
		t.Errorf("Expected set to contain 'b'")
	}
	if !s.Contains("a") {
		t.Errorf("Expected set to contain 'a'")
	}
	if s.Contains("c") {
		t.Errorf("Expected set to not contain 'c'")
	}

	s.Remove("a")
	if s.Contains("a") {
		t.Errorf("Expected set to not contain 'a'")
	}

	if s.IsEmpty() {
		t.Errorf("Expected set to not be empty")
	}

	all := s.GetAll()
	if len(all) != 1 || all[0] != "b" {
		t.Errorf("Expected set to contain only 'b', got %v", all)
	}

	s.Remove("b")
	if !s.IsEmpty() {
		t.Errorf("Expected set to be empty")
	}
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}

	s.Remove("nonexistent")
	if !s.IsEmpty() {
		t.Errorf("Expected set to be empty after removing nonexistent element")
	}
}
