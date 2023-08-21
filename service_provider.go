package godi

import (
	"fmt"
	"reflect"
)

const (
	singleton = iota
	provider
)

type serviceProvider struct {
	spType    int
	singleton any
	provide   interface{}
}

var spMap map[string]*serviceProvider = make(map[string]*serviceProvider)

func getKey[P any]() string {
	t := reflect.TypeOf((*P)(nil)).Elem()
	return t.PkgPath() + "." + t.Name()
}

func RegisterSingleton[P any](service P) {
	key := getKey[P]()
	sp := &serviceProvider{
		spType:    singleton,
		singleton: service,
	}
	spMap[key] = sp

}

func RegisterProvider[P any](provide func() (P, error)) {
	key := getKey[P]()
	sp := &serviceProvider{
		spType:  provider,
		provide: provide,
	}
	spMap[key] = sp
}

func GetService[P any]() (P, error) {
	key := getKey[P]()
	var zero P
	sp, ok := spMap[key]
	if !ok {
		return zero, fmt.Errorf("No provider found for type '%s'", key)
	}

	if sp.spType == singleton {
		return sp.singleton.(P), nil
	}

	provide := sp.provide.(func() (P, error))
	return provide()
}
