package stream

import "reflect"

func (s Stream) Error(err error) Stream {
	s.value = reflect.ValueOf(nil)
	s.elemType = nil
	s.len = 0
	s.slice = nil
	// TODO: maybe store multiple errors instead of overwriting
	s.err = err
	return s
}
