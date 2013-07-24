package lc_gospec

import (
	"github.com/orfjackal/gospec/src/gospec"
	"testing"
)

func TestAllSpecs(t *testing.T) {
	r := gospec.NewRunner()
	r.AddSpec(LoggregatorClientSpec)
	gospec.MainGoTest(r, t)
}
