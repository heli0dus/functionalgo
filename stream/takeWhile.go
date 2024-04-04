package stream

import (
	"fmt"
	"reflect"
)

func (s Stream) TakeWhile(f interface{}) Stream {
	if !s.Valid() {
		return s
	}
	fType := reflect.TypeOf(f)

	if fType.Kind() != reflect.Func {
		return s.Error(fmt.Errorf("TakeWhile expected function, but got %v", fType))
	}

	// return type check
	if !(fType.Out(0).Kind() == reflect.Bool &&
		(fType.NumOut() == 1 ||
			(fType.NumOut() == 2 &&
				fType.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem())))) {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("function used in TakeWhile must return bool or (bool, error)"))
	}

	// param type check
	if fType.NumIn() != 1 {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("function used in TakeWhile must receive 1 argument"))
	}
	if fType.In(0) != s.elemType {
		// TODO: better error message with actual types
		return s.Error(fmt.Errorf("type of argument 1 in function must be the same as element type of a stream"))
	}

	newLen := 0
	if fType.NumIn() == 1 {
		for i := 0; i < s.len; i++ {
			args := append([]reflect.Value{}, s.value.Index(i))
			newVals := reflect.ValueOf(f).Call(args)
			if newVals[0].Bool() {
				newLen++
			} else {
				break
			}
		}
	} else {
		for i := 0; i < s.len; i++ {
			args := append([]reflect.Value{}, s.value.Index(i))
			newVals := reflect.ValueOf(f).Call(args)
			if !newVals[1].IsNil() {
				return s.Error(newVals[1].Interface().(error))
			}
			if newVals[0].Bool() {
				newLen++
			} else {
				break
			}
		}
	}

	s.value = s.value.Slice(0, newLen)
	s.slice = s.value.Interface()

	s.len = newLen

	return s
}
