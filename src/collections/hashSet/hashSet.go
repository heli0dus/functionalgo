package hashSet

type void struct{}

var member void

type HashSet[T comparable] map[T]void

func (set *HashSet[T]) add(x T) {
	(*set)[x] = member
}

func (set *HashSet[T]) remove(x T) bool {
	if _, ok := (*set)[x]; ok {
		delete((*set), x)
		return true
	}
	return false
}

func (set *HashSet[T]) size() int {
	return len(*set)
}

func fmap[T comparable, V comparable](set HashSet[T], f func(T) V) HashSet[V] {
	res := make(map[V]void)
	for k := range set {
		res[f(k)] = member
	}

	return res
}
