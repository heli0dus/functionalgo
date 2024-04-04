package stream

import (
	"fmt"
	"reflect"
)

func (s Stream) All(f interface{}) Stream {
	// TODO: replace with s.Valid()
	if s.err != nil {
		return s
	}
	fType := reflect.TypeOf(f)
	// return type check
	if !(fType.Out(0).Kind() == reflect.Bool && (fType.NumOut() == 1 || (fType.NumOut() == 2 && fType.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem())))) {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("function used in all must return bool or (bool, error)"))
	}

	// param type check
	if fType.NumIn() != 1 {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("function used in all must receive 1 argument"))
	}
	if fType.In(0) != s.elemType {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("type of argument 1 in function must be the same as element type of a stream"))
	}

	return s.Fmap(f).Reduce(func(a bool, b bool) bool {
		return a && b
	}, true)
}
