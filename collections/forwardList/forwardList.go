package forwardList

// Node представляет собой узел в списке.
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// ForwardList представляет собой односвязный список.
type ForwardList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

// NewForwardList создает новый пустой список.
func NewForwardList[T comparable]() *ForwardList[T] {
	return &ForwardList[T]{}
}

// Push добавляет новый элемент в конец списка.
func (fl *ForwardList[T]) Push(value T) {
	if fl.Head == nil {
		fl.Head = &Node[T]{Value: value}
		fl.Tail = fl.Head
		return
	}
	fl.Tail.Next = &Node[T]{Value: value}
	fl.Tail = fl.Tail.Next
}

// Find ищет первый элемент со значением, равным указанному, и возвращает его.
// Если такой элемент не найден, возвращает nil.
func (fl *ForwardList[T]) Find(value T) *Node[T] {
	for current := fl.Head; current != nil; current = current.Next {
		if current.Value == value {
			return current
		}
	}
	return nil
}

// Delete удаляет первый элемент со значением, равным указанному.
func (fl *ForwardList[T]) Delete(value T) {
	if fl.Head == nil {
		return
	}
	if fl.Head.Value == value {
		fl.Head = fl.Head.Next
		return
	}
	for current := fl.Head; current.Next != nil; current = current.Next {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return
		}
	}
}

func Fmap[T comparable, V comparable](fl *ForwardList[T], f func(T) V) *ForwardList[V] {
	newList := NewForwardList[V]()

	for current := fl.Head; current != nil; current = current.Next {
		transformedValue := f(current.Value)
		newList.Push(transformedValue)
	}

	return newList
}

func Reduce[T comparable, V comparable](fl *ForwardList[T], initial V, f func(V, T) V) V {
	result := initial

	for current := fl.Head; current != nil; current = current.Next {
		result = f(result, current.Value)
	}

	return result
}

func fromSlice[T comparable](s []T) *ForwardList[T] {
	list := NewForwardList[T]()
	for i := 0; i < len(s); i++ {
		list.Push(s[i])
	}
	return list
}

func (fl *ForwardList[T]) toSlice() []T {
	var slice []T
	for current := fl.Head; current != nil; current = current.Next {
		slice = append(slice, current.Value)
	}
	return slice
}
