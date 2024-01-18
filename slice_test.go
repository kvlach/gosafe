package gosafe

import (
	"testing"
)

func TestSlice(t *testing.T) {
	s := Slice[int]{}

	if l := s.Len(); l != 0 {
		t.Logf("Expected length to be 0, got %d", l)
		t.Fail()
	}

	const v1 = 123
	const v2 = 456
	const v3 = 789

	s.Append(v1, v2)
	if l := s.Len(); l != 2 {
		t.Logf("Expected length to be 2, got %d", l)
		t.Fail()
	}
	if v := s.Get(1); v != v2 {
		t.Logf("Expected element with value %d, got %d", v2, v)
		t.Fail()
	}
	if !s.In(v1) {
		t.Logf("Expected to find element in slice with value %d, but didn't", v1)
		t.Fail()
	}

	s.Set(1, v3)
	if v := s.Get(1); v != v3 {
		t.Logf("Expected element with value %d, got %d", v3, v)
		t.Fail()
	}

	s.DeleteStable(0)
	s.DeleteUnstable(0)
	if l := s.Len(); l != 0 {
		t.Logf("Expected length to be 0 but got %d", l)
		t.Fail()
	}
}
