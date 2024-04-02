package stream

import (
	"fmt"
	"reflect"
)

type Stream struct {
	slice    interface{}
	value    reflect.Value
	elemType reflect.Type
	len      int
	err      error
}

type Streamable[T any] interface {
	ToSlice() []T
}

func AsStream[T any](slice []T) Stream {
	s := Stream{}
	s.value = reflect.ValueOf(slice)
	s.elemType = reflect.TypeOf(slice).Elem()
	s.len = len(slice)
	s.slice = slice
	return s
}

func AsSlice[T any](s Stream) ([]T, error) {
	if s.err != nil {
		return nil, s.err
	}
	if reflect.TypeOf((*T)(nil)).Elem() != s.elemType {
		return nil, fmt.Errorf("cannot convert stream of %v to []%v", s.elemType, reflect.TypeOf((*T)(nil)).Elem())
	}
	return s.slice.([]T), nil
}
