package stream

import (
	"fmt"
	"reflect"
)

func (s Stream) Filter(f interface{}) Stream {
	if !s.Valid() {
		return s
	}

	fType := reflect.TypeOf(f)

	if fType.Kind() != reflect.Func {
		return s.Error(fmt.Errorf("in Filter expected function, but got %v", fType))
	}

	// return type check
	if !(fType.Out(0).Kind() == reflect.Bool && (fType.NumOut() == 1 || (fType.NumOut() == 2 && fType.Out(1).Implements(reflect.TypeFor[error]())))) {
		return s.Error(fmt.Errorf("function used in Filter must return bool or (bool, error)"))
	}

	// param type check
	if fType.NumIn() != 1 {
		return s.Error(fmt.Errorf("function used in Filter must receive 1 argument"))
	}
	if fType.In(0) != s.elemType {
		return s.Error(fmt.Errorf("type of argument 1 in function must be the same as element type of a stream"))
	}

	res := reflect.MakeSlice(reflect.SliceOf(s.elemType), s.len, s.len)
	resI := 0
	if fType.NumOut() == 1 {
		for i := 0; i < s.len; i++ {
			newVals := reflect.ValueOf(f).Call([]reflect.Value{s.value.Index(i)})
			if newVals[0].Bool() {
				res.Index(resI).Set(s.value.Index(i))
				resI++
			}
		}
	} else {
		for i := 0; i < s.len; i++ {
			newVals := reflect.ValueOf(f).Call([]reflect.Value{s.value.Index(i)})
			if !newVals[1].IsNil() {
				return s.Error(newVals[1].Interface().(error))
			}
			if newVals[0].Bool() {
				res.Index(resI).Set(s.value.Index(i))
				resI++
			}
		}
	}
	res.Slice(0, resI)
	s.len = resI
	s.value = res
	s.slice = res.Interface()
	return s
}
