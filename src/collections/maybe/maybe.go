package maybe

import "fmt"

type Maybe[T any] struct {
	empty bool
	data  T
}

func (m *Maybe[T]) Empty() bool {
	return (*m).empty
}

func (m *Maybe[T]) Value() *T {
	return &m.data
}

func MaybeOf[T any](x T) Maybe[T] {
	res := Maybe[T]{}
	res.empty = false
	res.data = x
	return res
}

func FromMaybe[T any, V any](deflt V, f func(T) V, m *Maybe[T]) V {
	if m.Empty() {
		return deflt
	}
	return f(*m.Value())
}

func EmptyMaybe[T any]() Maybe[T] {
	res := Maybe[T]{}
	res.empty = true
	return res
}

func Fmap[T any, V any](m Maybe[T], f func(T) V) Maybe[V] {
	if m.Empty() {
		res := new(Maybe[V])
		res.empty = true
		return *res
	}
	res := new(Maybe[V])
	res.empty = false
	res.data = f(m.data)
	return *res
}

func ToSlice[T any](m Maybe[T]) []T {
	if m.Empty() {
		return make([]T, 0)
	}
	return []T{m.data}
}

func FromSlice[T any](a []T) (Maybe[T], error) {
	switch len(a) {
	case 0:
		return EmptyMaybe[T](), nil
	case 1:
		return MaybeOf(a[0]), nil
	default:
		return Maybe[T]{}, fmt.Errorf("error: too much elementa to construct Maybe")
	}

}
