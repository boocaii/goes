package maps

import "reflect"

func NewMap(in ...interface{}) interface{} {
	var out interface{}

	typeOfIn := reflect.TypeOf(in)
	switch typeOfIn.Kind() {
	case reflect.Array:

	}

	return out
}
