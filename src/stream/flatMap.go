package stream

import (
	"fmt"
	"reflect"
)

func (s Stream) FlatMap(f interface{}) Stream {
	if !s.Valid() {
		return s
	}
	fType := reflect.TypeOf(f)

	// return type check
	// TODO: rewrite with lines of sane length
	if !((fType.NumOut() == 1 && fType.Out(0).Kind() == reflect.Slice) || (fType.NumOut() == 2 && fType.Out(1).Kind() == reflect.Slice && fType.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem()))) {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("function used in FlatMap must return slice or (slice, error)"))
	}

	// param type check
	if fType.NumIn() != 1 {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("function used in FlatMap must receive 1 argument"))
	}
	if fType.In(0) != s.elemType {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("type of argument 1 in function must be the same as element type of a stream"))
	}

	res := reflect.MakeSlice(reflect.SliceOf(fType.Out(0).Elem()), 0, 0)
	for i := 0; i < s.len; i++ {
		args := append([]reflect.Value{}, s.value.Index(i))
		newVals := reflect.ValueOf(f).Call(args)
		if len(newVals) == 2 && !newVals[1].IsNil() {
			return s.Error(newVals[1].Interface().(error))
		}
		res = reflect.AppendSlice(res, newVals[0])
	}
	s.elemType = res.Type().Elem()
	s.value = res
	s.slice = res.Interface()
	s.len = res.Len()
	return s
}
