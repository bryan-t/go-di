package godi

import (
	"fmt"
	"reflect"
)

const (
	singleton = iota
	factory
)

type serviceProvider struct {
	spType    int
	singleton any
	provide   interface{}
}

var spMap map[string]*serviceProvider = make(map[string]*serviceProvider)

func RegisterSingleton[P any](service P) {
	t := reflect.TypeOf((*P)(nil)).Elem()
	key := t.PkgPath() + "." + t.Name()

	sp := &serviceProvider{
		spType:    singleton,
		singleton: service,
	}
	spMap[key] = sp

}

func RegisterProvider[P any](provide func() (P, error)) {
	t := reflect.TypeOf((*P)(nil)).Elem()
	key := t.PkgPath() + "." + t.Name()

	sp := &serviceProvider{
		spType:  singleton,
		provide: provide,
	}
	spMap[key] = sp
}

func GetService[P any]() (P, error) {
	var zero P
	t := reflect.TypeOf((*P)(nil)).Elem()
	key := t.PkgPath() + "." + t.Name()
	sp, ok := spMap[key]
	if !ok {
		return zero, fmt.Errorf("No provider found for type '%s'", t.Name())
	}

	if sp.spType == singleton {
		return sp.singleton.(P), nil
	}
	provide := sp.provide.(func() (P, error))
	return provide()
}
