package interfaces

type Functor[T any, V any] interface {
	fmap(func(T) V) Functor[V, V]
}

type Endo[T any] Functor[T, T]
