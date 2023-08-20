package godi

import (
	"reflect"
)

var sp map[string]interface{} = make(map[string]interface{})

func RegisterSingleton[P any](service P) {
	t := reflect.TypeOf((*P)(nil)).Elem()
	key := t.PkgPath() + "." + t.Name()
	sp[key] = service
}
func GetService[P any]() P {
	var zero_p P

	return zero_p
}
