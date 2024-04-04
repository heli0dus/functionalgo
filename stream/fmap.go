package stream

import (
	"fmt"
	"reflect"
)

func (s Stream) Fmap(f interface{}) Stream {
	if !s.Valid() {
		return s
	}

	fType := reflect.TypeOf(f)

	if fType.Kind() != reflect.Func {
		return s.Error(fmt.Errorf("in Fmap expected function, but got %v", fType))
	}

	// return type check
	if !(fType.NumOut() == 1 || (fType.NumOut() == 2 && fType.Out(1).Implements(reflect.TypeFor[error]()))) {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("function used in fmap must return value or (value, error)"))
	}

	// param type check
	if fType.NumIn() != 1 {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("function used in fmap must receive 1 argument"))
	}
	if fType.In(0) != s.elemType {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("type of argument 1 in function must be the same as element type of a stream"))
	}

	res := reflect.MakeSlice(reflect.SliceOf(fType.Out(0)), s.len, s.len)
	if fType.NumOut() == 1 {
		for i := 0; i < s.len; i++ {
			newVals := reflect.ValueOf(f).Call([]reflect.Value{s.value.Index(i)})
			res.Index(i).Set(newVals[0])
		}
	} else {
		for i := 0; i < s.len; i++ {
			newVals := reflect.ValueOf(f).Call([]reflect.Value{s.value.Index(i)})
			if !newVals[1].IsNil() {
				return s.Error(newVals[1].Interface().(error))
			}
			res.Index(i).Set(newVals[0])
		}
	}
	s.elemType = res.Type().Elem()
	s.value = res
	s.slice = res.Interface()
	return s
}
