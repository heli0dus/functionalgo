package hashSet

type void struct{}

var member void

type HashSet[T comparable] map[T]void

func (set *HashSet[T]) Add(x T) {
	(*set)[x] = member
}

func (set *HashSet[T]) Remove(x T) bool {
	if _, ok := (*set)[x]; ok {
		delete((*set), x)
		return true
	}
	return false
}

func (set *HashSet[T]) Size() int {
	return len(*set)
}

func (set *HashSet[T]) Contains(x T) bool {
	_, ok := (*set)[x]
	return ok
}

func New[T comparable]() HashSet[T] {
	return make(HashSet[T])
}

func Fmap[T comparable, V comparable](set HashSet[T], f func(T) V) HashSet[V] {
	res := make(map[V]void)
	for k := range set {
		res[f(k)] = member
	}

	return res
}

func ToSlice[T comparable](set HashSet[T]) []T {
	res := make([]T, 0, len(set))
	for k := range set {
		res = append(res, k)
	}
	return res
}

func FromSlice[T comparable](arr []T) HashSet[T] {
	res := HashSet[T]{}
	for _, val := range arr {
		res.Add(val)
	}
	return res
}
