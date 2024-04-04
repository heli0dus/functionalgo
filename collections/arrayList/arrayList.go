package arraylist

type ArrayList[T any] []T

// Finds element by index. Panics if index is out of bounds
func (arr *ArrayList[T]) At(i int) T {
	return (*arr)[i]
}

// Appends element to the end of list
func (arr *ArrayList[T]) Append(elem T) {
	(*arr) = append((*arr), elem)
}

// Returns size of array
func (arr *ArrayList[T]) Size() int {
	return len(*arr)
}

// Converts ArrayList to slice
func (arr *ArrayList[T]) ToSlice() []T {
	return *arr
}

func New[T any]() ArrayList[T] {
	return make(ArrayList[T], 0)
}

// Just casts type. Added for consistency between collections
func FromSlice[T any](arr []T) ArrayList[T] {
	return arr
}

// Maps function to elements of list
func Fmap[T any, V any](arr ArrayList[T], f func(T) V) ArrayList[V] {
	res := make([]V, 0, len(arr))
	for _, val := range arr {
		res = append(res, f(val))
	}

	return res
}

// Folds list with given operation and started element into one element
func Reduce[A any, B any](arr ArrayList[A], ini B, f func(B, A) B) B {
	res := ini
	for _, val := range arr {
		res = f(res, val)
	}
	return res
}
