package stream

import (
	"fmt"
	"reflect"
)

func (s Stream) Reduce(f interface{}, init interface{}) Stream {
	if !s.Valid() {
		return s
	}

	fType := reflect.TypeOf(f)

	if fType.Kind() != reflect.Func {
		return s.Error(fmt.Errorf("in Reduce expected function, but got %v", fType))
	}

	if !(fType.NumOut() == 1 || (fType.NumOut() == 2 && fType.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem()))) {
		return s.Error(fmt.Errorf("function used in Reduce must return value or (value, error)"))
	}

	if fType.NumIn() != 2 {
		return s.Error(fmt.Errorf("function used in reduce must receive 2 arguments"))
	}

	if fType.Out(0) != fType.In(0) {
		return s.Error(fmt.Errorf("output of function must have the same type as its first argument"))
	}

	if fType.In(0) != reflect.TypeOf(init) {
		return s.Error(fmt.Errorf("init must have the same type as the first argument of function"))
	}

	if fType.In(1) != s.elemType {
		return s.Error(fmt.Errorf("type of argument 2 in function must be the same as element type of a stream"))
	}

	res := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(f).Out(0)), 0, 1)
	res = reflect.Append(res, reflect.ValueOf(init))
	for i := 0; i < s.len; i++ {
		args := append([]reflect.Value{}, res.Index(0), s.value.Index(i))
		newVals := reflect.ValueOf(f).Call(args)
		if len(newVals) == 2 && !newVals[1].IsNil() {
			return s.Error(newVals[1].Interface().(error))
		}
		res.Index(0).Set(newVals[0])
	}
	s.elemType = res.Type().Elem()
	s.value = res
	s.slice = res.Interface()
	s.len = 1
	return s
}
