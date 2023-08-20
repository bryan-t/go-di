package godi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tester interface {
	test()
}
type testProvider struct{}

func (t *testProvider) test() {}

func TestRegisterSingleton(t *testing.T) {
	impl := &testProvider{}
	RegisterSingleton[tester](impl)
	assert.Equal(t, impl, sp["github.com/bryan-t/go-di.git.tester"], "Interface not found")

	x := 5
	RegisterSingleton[int](x)
	assert.Equal(t, x, sp[".int"], "Int not found")
}
