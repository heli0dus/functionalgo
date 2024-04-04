package stream

import "testing"

func TestTakeWhileConstTrue(t *testing.T) {
	slice := []int{1, 2, 3}
	s := AsStream(slice)
	s = s.TakeWhile(func(x int) bool { return true })
	newSlice, err := AsSlice[int](s)

	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if len(newSlice) != 3 {
		t.Errorf("expected one return value, but got: %v", len(newSlice))
	}
}

func TestTakeWhileConstFalse(t *testing.T) {
	slice := []int{1, 2, 3}
	s := AsStream(slice)
	s = s.TakeWhile(func(x int) bool { return false })
	newSlice, err := AsSlice[int](s)

	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if len(newSlice) != 0 {
		t.Errorf("expected one return value, but got: %v", len(newSlice))
	}
}

func TestTakeWhileRegular(t *testing.T) {
	slice := []int{1, 2, 3}
	s := AsStream(slice)
	s = s.TakeWhile(func(x int) bool { return x%2 != 0 })
	newSlice, err := AsSlice[int](s)

	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if len(newSlice) != 1 {
		t.Errorf("expected one return value, but got: %v", len(newSlice))
	} else if newSlice[0] != 1 {
		t.Errorf("unexpected value of first element of slice. expected 1, but got %v", newSlice[0])
	}
}
