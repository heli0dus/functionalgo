package stream

import (
	"fmt"
	"reflect"
)

func (s Stream) All(f interface{}) Stream {
	if !s.Valid() {
		return s
	}

	fType := reflect.TypeOf(f)

	if fType.Kind() != reflect.Func {
		return s.Error(fmt.Errorf("in All expected function, but got %v", fType))
	}

	// return type check
	if !(fType.Out(0).Kind() == reflect.Bool && (fType.NumOut() == 1 || (fType.NumOut() == 2 && fType.Out(1).Implements(reflect.TypeFor[error]())))) {
		return s.Error(fmt.Errorf("function used in All must return bool or (bool, error)"))
	}

	// param type check
	if fType.NumIn() != 1 {
		return s.Error(fmt.Errorf("function used in All must receive 1 argument"))
	}
	if fType.In(0) != s.elemType {
		return s.Error(fmt.Errorf("type of argument 1 in function must be the same as element type of a stream"))
	}

	return s.Fmap(f).Reduce(func(a bool, b bool) bool {
		return a && b
	}, true)
}
