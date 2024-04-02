package functools

func Compose[T any, V any, R any](f func(T) V, g func(V) R) func(T) R {
	return func(x T) R {
		return g(f(x))
	}
}

func On[A any, B any, C any](f func(B, B) C, g func(A) B) func(A, A) C {
	return func(x, y A) C {
		return f(g(x), g(y))
	}
}

func Curry[A any, B any, C any](f func(A, B) C) func(A) func(B) C {
	return func(x A) func(B) C {
		return func(y B) C {
			return f(x, y)
		}
	}
}
