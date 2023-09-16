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
	assert.Equal(t, impl, spMap["github.com/bryan-t/godi.git.tester"].singleton, "Interface not found")

	x := 5
	RegisterSingleton[int](x)
	assert.Equal(t, x, spMap[".int"].singleton, "Int not found")
}

func TestRegisterProvider(t *testing.T) {
	impl := &testProvider{}
	RegisterProvider[tester](func() (tester, error) { return impl, nil })

	registeredImpl, _ := spMap["github.com/bryan-t/godi.git.tester"].provide.(func() (tester, error))()
	assert.Equal(t, impl, registeredImpl, "Got a different impl from provide")
}

func TestGetService(t *testing.T) {

	// Get singleton
	//
	impl := &testProvider{}
	RegisterSingleton[tester](impl)
	registeredImpl, _ := GetService[tester]()
	assert.Equal(t, impl, registeredImpl, "registeredImpl not the same.")

	x := 5
	RegisterSingleton[int](x)
	registerdInt, _ := GetService[int]()
	assert.Equal(t, x, registerdInt, "registerdInt not the same.")

	// Get provide
	//
	RegisterProvider[tester](func() (tester, error) { return impl, nil })
	registeredImpl, _ = GetService[tester]()
	assert.Equal(t, impl, registeredImpl, "registeredImpl not the same using provide.")
}

func TestLoad(t *testing.T) {
	// create config
	//
	config := Config{
		Singletons: []Singleton{
			Singleton{
				InterfacePkg: "github.com/bryan-t/godi",
				Interface:    "tester",
			},
		},
	}

	err := Load(config)
	assert.Nil(t, err)
}
