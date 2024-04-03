package stream

import (
	"fmt"
	"reflect"
)

func (s Stream) DropWhile(f interface{}) Stream {
	if !s.Valid() {
		return s
	}
	funcType := reflect.TypeOf(f)

	if !(funcType.Out(0).Kind() == reflect.Bool &&
		(funcType.NumOut() == 1 ||
			(funcType.NumOut() == 2 &&
				funcType.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem())))) {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("function used in `DropWhile` must return bool or (bool, error)"))
	}

	if funcType.NumIn() != 1 {
		return s.Error(fmt.Errorf("function used in `DropWhile` must receive 1 argument"))
	}
	if funcType.In(0) != s.elemType {
		return s.Error(fmt.Errorf("type of argument 1 in function must be the same as element type of a stream"))
	}

	startFrom := 0
	for ; startFrom < s.len; startFrom++ {
		args := append([]reflect.Value{}, s.value.Index(startFrom))
		check := reflect.ValueOf(f).Call(args)
		if len(check) == 2 && !check[1].IsNil() {
			return s.Error(check[1].Interface().(error))
		}
		if !check[0].Bool() {
			break
		}
	}
	// now we must get all remaining elements
	newLen := s.len - startFrom
	s.value = s.value.Slice(startFrom, s.len)
	s.len = newLen
	s.slice = s.value.Interface()

	return s
}
