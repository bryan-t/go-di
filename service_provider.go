package godi

import (
	"fmt"
	"go/importer"
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

func Load(config Config) error {

	for _, v := range config.Singletons {
		interfacePkg, err := importer.Default().Import(v.InterfacePkg)
		if err != nil {
			return err
		}
		interfacePkgScope := interfacePkg.Scope()
		interfaceObject := interfacePkgScope.Lookup(v.Interface)
		if interfaceObject == nil {
			return fmt.Errorf("Did not find interface '%s'", v.Interface)
		}
		fmt.Printf("Got interface %+v \n", interfaceObject)

	}
	return nil
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
