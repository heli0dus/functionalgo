package collections

// Node представляет собой узел в списке.
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// ForwardList представляет собой односвязный список.
type ForwardList[T comparable] struct {
	Head *Node[T]
}

// NewForwardList создает новый пустой список.
func NewForwardList[T comparable]() *ForwardList[T] {
	return &ForwardList[T]{}
}

// PushFront добавляет новый элемент в начало списка.
func (fl *ForwardList[T]) PushFront(value T) {
	newNode := &Node[T]{Value: value, Next: fl.Head}
	fl.Head = newNode
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
		newList.PushFront(transformedValue)
	}

	return newList
}

func Reduce[T comparable, V comparable](fl *ForwardList[T], initial V, f func(V, T) V) V {
	result := initial

	for current := fl.Head; current != nil; current = current.Next {
		result = f(result, current.Value)
	}

	// Возвращаем итоговое значение
	return result
}
