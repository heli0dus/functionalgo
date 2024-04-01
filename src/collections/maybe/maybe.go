package maybe

type Maybe[T any] struct {
	empty bool
	data  T
}

func (m *Maybe[T]) Empty() bool {
	return (*m).empty
}

func fmap[T any, V any](m Maybe[T], f func(T) V) Maybe[V] {
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
