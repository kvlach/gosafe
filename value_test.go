package gosafe

import (
	"testing"
)

func TestValue(t *testing.T) {
	n := Value[int]{}

	if val := n.Get(); val != 0 {
		t.Logf("expected value to be 0 got %d instead", val)
		t.Fail()
	}

	const test = 123

	n.Set(test)

	if val := n.Get(); val != test {
		t.Logf("expected value to be %d got %d instead", test, val)
		t.Fail()
	}
}
