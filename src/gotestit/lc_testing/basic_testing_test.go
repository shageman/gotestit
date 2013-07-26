package lc_testing

import "testing"

func TestSomething(t *testing.T) {
	if 1 == 2 {
		t.Error("1 shouldn't equal 2")
	}
}
