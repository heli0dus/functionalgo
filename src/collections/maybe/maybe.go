package maybe

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
