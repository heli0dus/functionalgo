package stream

import "reflect"

type Error struct{}

type Stream struct {
	slice    interface{}
	value    reflect.Value
	elemType reflect.Type
	len      int
	err      *Error
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

func AsSlice[T any](s Stream) []T {
	// TODO: add type check
	return s.slice.([]T)
}
