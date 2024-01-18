package gosafe

import (
	"testing"
)

func TestMap(t *testing.T) {
	m := Map[int, string]{}

	const key = 123
	const value = "abc"

	if v, ok := m.Get(key); ok {
		t.Logf("Found key %d with value '%s' even though it was never set", key, v)
		t.Fail()
	}

	m.Set(key, value)
	v, ok := m.Get(key)
	if !ok {
		t.Logf("Didn't find key %d", key)
		t.Fail()
	}
	if v != value {
		t.Logf("Expected value '%s' got '%s' instead", value, v)
		t.Fail()
	}

	// delete non-existant key, this shouldn't fail
	m.Delete(456)

	m.Delete(key)
	if _, ok := m.Get(key); ok {
		t.Logf("Found key %d when it should have been deleted", key)
		t.Fail()
	}
}
