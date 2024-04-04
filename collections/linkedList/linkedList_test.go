package linkedlist

import "testing"

func TestCreate(t *testing.T) {
	list := New[int]()
	if list.Size() != 0 {
		t.Errorf("list created with non-empty state")
	}

	if list.Front() != nil {
		t.Errorf("Front on empty list returns non-nil")
	}

	if list.Back() != nil {
		t.Errorf("Back on empty list returns non-nil")
	}

}

func TestInsert(t *testing.T) {
	list := New[int]()
	list.Insert(1)
	list.Insert(2)
	it := list.Back()
	list.Insert(5)

	it = list.InsertAfter(4, it)
	list.InsertBefore(3, it)

	i := 1
	for it := list.Front(); it != nil; it = it.Next() {
		if it.Value != i {
			t.Errorf("Insertion failed on element %d", i)
			return
		}
		i += 1
	}
}

func TestSlice(t *testing.T) {
	arr := []string{
		"one",
		"two",
		"three",
	}

	list := FromSlice(arr)

	i := 0
	for it := list.Front(); it != nil; it = it.Next() {
		if it.Value != arr[i] {
			t.Errorf("Conversion from slice failed at element %d", i)
		}
		i += 1
	}
	newArr := list.ToSlice()
	i = 0
	for ; i != len(arr); i += 1 {
		if newArr[i] != arr[i] {
			t.Errorf("Conversion to slice failed at element %d", i)
		}
	}
}
