package stream

import (
	"fmt"
	"reflect"
)

func (s Stream) Filter(f interface{}) Stream {
	// TODO: replace with s.Valid()
	if s.err != nil {
		return s
	}
	fType := reflect.TypeOf(f)
	// return type check
	if !(fType.Out(0).Kind() == reflect.Bool && (fType.NumOut() == 1 || (fType.NumOut() == 2 && fType.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem())))) {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("function used in filter must return bool or (bool, error)"))
	}

	// param type check
	if fType.NumIn() != 1 {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("function used in filter must receive 1 argument"))
	}
	if fType.In(0) != s.elemType {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("type of argument 1 in function must be the same as element type of a stream"))
	}

	res := reflect.MakeSlice(reflect.SliceOf(s.elemType), 0, 0)
	newLen := 0
	for i := 0; i < s.len; i++ {
		args := append([]reflect.Value{}, s.value.Index(i))
		newVals := reflect.ValueOf(f).Call(args)
		if len(newVals) == 2 && !newVals[1].IsNil() {
			return s.Error(newVals[1].Interface().(error))
		}
		if newVals[0].Bool() {
			res = reflect.Append(res, s.value.Index(i))
			newLen++
		}
	}
	s.len = newLen
	s.elemType = res.Type().Elem()
	s.value = res
	s.slice = res.Interface()
	return s
}
