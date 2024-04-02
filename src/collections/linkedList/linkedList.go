package linkedlist

type Element[T any] struct {
	Value      T
	next, prev *Element[T]

	// original list to get essential values like size, begin and end
	list *LinkedList[T]
}

func (self *Element[T]) Next() *Element[T] {
	if self.next != &self.list.root {
		return self.next
	}
	return nil
}

func (self *Element[T]) Prev() *Element[T] {
	if self.prev != &self.list.root {
		return self.prev
	}
	return nil
}

type LinkedList[T any] struct {
	// root is a dummy element to create ring from nodes
	root Element[T]
	size int
}

func New[T any]() *LinkedList[T] {
	return new(LinkedList[T]).Init()
}

func (self *LinkedList[T]) Init() *LinkedList[T] {
	self.root.next = &self.root
	self.root.prev = &self.root
	self.size = 0
	self.root.list = self
	return self
}

func (self *LinkedList[T]) Size() int {
	return self.size
}

func (self *LinkedList[T]) Back() *Element[T] {
	if self.size == 0 {
		return nil
	}
	return self.root.prev
}

func (self *LinkedList[T]) Front() *Element[T] {
	if self.size == 0 {
		return nil
	}
	return self.root.next
}

// inserts value after specified element
// Returns: pointer to inserted element
func (self *LinkedList[T]) InsertAfter(elem T, e *Element[T]) *Element[T] {
	var newNode = new(Element[T])
	newNode.Value = elem
	newNode.prev = e
	newNode.next = e.next
	newNode.list = self

	e.next.prev = newNode
	e.next = newNode

	self.size += 1

	return newNode
}

func (self *LinkedList[T]) InsertBefore(elem T, e *Element[T]) *Element[T] {
	return self.InsertAfter(elem, e.prev)
}

func (self *LinkedList[T]) Insert(elem T) *Element[T] {
	return self.InsertAfter(elem, self.root.prev)
}

func (self *LinkedList[T]) ToSlice() []T {
	var result []T
	for it := self.root.next; it != &self.root; it = it.next {
		result = append(result, it.Value)
	}
	return result
}

func FromSlice[T any](slice []T) *LinkedList[T] {
	list := New[T]()
	for i := 0; i < len(slice); i++ {
		list.Insert(slice[i])
	}
	return list
}
