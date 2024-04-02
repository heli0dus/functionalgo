package arraylist

type ArrayList[T any] []T

func (arr *ArrayList[T]) At(i int) T {
	return (*arr)[i]
}

func (arr *ArrayList[T]) Append(elem T) {
	(*arr) = append((*arr), elem)
}

func (arr *ArrayList[T]) Size() int {
	return len(*arr)
}

func (arr *ArrayList[T]) ToSlice() []T {
	return *arr
}

func Fmap[T any, V any](arr ArrayList[T], f func(T) V) ArrayList[V] {
	res := make([]V, 0, len(arr))
	for _, val := range arr {
		res = append(res, f(val))
	}

	return res
}

func Reduce[A any, B any](arr ArrayList[A], ini B, f func(B, A) B) B {
	res := ini
	for _, val := range arr {
		res = f(res, val)
	}
	return res
}

// func (arr *ArrayList[T, V]) fmap(f func(T) V) ArrayList[V, V] {
// res := make([]V, 0, len(*arr))
// for _, val := range *arr {
// 	res = append(res, f(val))
// }

// return res
// }
