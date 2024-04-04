package stream

import (
	"slices"
	"testing"
)

func TestFilterConstTrue(t *testing.T) {
	slice := []int{1, 2, 3}
	s := AsStream(slice)
	s = s.Filter(func(x int) bool { return true })
	newSlice, err := AsSlice[int](s)

	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if !slices.Equal(slice, newSlice) {
		t.Errorf("expected %v, but got: %v", slice, newSlice)
	}
}

func TestFilterConstFalse(t *testing.T) {
	slice := []int{1, 2, 3}
	s := AsStream(slice)
	s = s.Filter(func(x int) bool { return false })
	newSlice, err := AsSlice[int](s)

	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if len(newSlice) != 0 {
		t.Errorf("expected empty result, but got: %v", newSlice)
	}
}

func TestFilterRegular(t *testing.T) {
	slice := []int{1, 2, 3}
	s := AsStream(slice)
	s = s.Filter(func(x int) bool { return x%2 != 0 })
	newSlice, err := AsSlice[int](s)

	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if len(newSlice) != 2 {
		t.Errorf("expected two return values, but got: %v", len(newSlice))
	} else if !slices.Equal(newSlice, []int{1, 3}) {
		t.Errorf("expected %v, but got %v", []int{1, 3}, newSlice)
	}
}
