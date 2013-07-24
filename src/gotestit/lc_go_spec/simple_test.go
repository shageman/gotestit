package lc_go_spec

import (
	. "github.com/bmatsuo/go-spec/spec"
	"testing"
)

func TestSimpleStuff(t *testing.T) {
	s := NewSpecTest(t)

	s.Describe("Value", func() {
		s.It("can be a native Go type", func() {
			s.Spec(1, Should, Equal, 1)
			s.Spec("abc", Should, Not, Equal, 1+3i)
			s.Spec(false, Should, Not, Equal, true)
		})
		s.It("can't fail!!!", func() {
			s.Spec(1, Should, Equal, 2)
			s.Spec("abc", Should, Equal, 1+3i)
			s.Spec(false, Should, Equal, true)
		})
	})
}
