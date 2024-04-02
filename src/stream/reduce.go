package stream

import (
	"fmt"
	"reflect"
)

func (s Stream) Reduce(f interface{}, init interface{}) Stream {
	// TODO: replace with s.Valid()
	if s.err != nil {
		return s
	}

	fType := reflect.TypeOf(f)
	if !(fType.NumOut() == 1 || (fType.NumOut() == 2 && fType.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem()))) {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("function used in reduce must return value or (value, error)"))
	}

	if fType.NumIn() != 2 {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("function used in reduce must receive 2 arguments"))
	}

	// TODO: this check can be less strict
	if fType.Out(0) != fType.In(0) {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("type mismatch in reduce"))
	}

	// TODO: this check can be less strict
	if fType.In(0) != reflect.TypeOf(init) {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("type mismatch in reduce"))
	}

	// TODO: this check can be less strict
	if fType.In(1) != s.elemType {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("type mismatch in reduce"))
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
