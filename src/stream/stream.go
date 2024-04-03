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
	// TODO: replace with s.Valid()
	if s.err != nil {
		return nil, s.err
	}
	if reflect.TypeOf((*T)(nil)).Elem() != s.elemType {
		return nil, fmt.Errorf("cannot convert stream of %v to []%v", s.elemType, reflect.TypeOf((*T)(nil)).Elem())
	}
	return s.slice.([]T), nil
}

func As[T any](s Stream) (T, error) {
	// TODO: replace with s.Valid()
	if s.err != nil {
		var t T
		return t, s.err
	}
	if reflect.TypeOf((*T)(nil)).Elem() != s.elemType {
		var t T
		return t, fmt.Errorf("cannot convert %v to []%v", s.elemType, reflect.TypeOf((*T)(nil)).Elem())
	}
	if s.len != 1 {
		var t T
		return t, fmt.Errorf("expected Stream with exactly 1 element")
	}
	return s.slice.([]T)[0], nil
}

func AsOr[T any](s Stream, def T) T {
	// TODO: replace with s.Valid()
	if s.err != nil {
		return def
	}
	if reflect.TypeOf((*T)(nil)).Elem() != s.elemType {
		return def
	}
	if s.len != 1 {
		return def
	}
	return s.slice.([]T)[0]

func (s Stream) Valid() bool {
	return s.err != nil

}
