package stream

import (
	"testing"
)

func TestTakeRegular(t *testing.T) {
	slice := []int{1, 2, 3}
	s := AsStream(slice)
	s = s.Take(2)
	newSlice, err := AsSlice[int](s)

	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if len(newSlice) != 2 {
		t.Errorf("expected one return value, but got: %v", len(newSlice))
	}
}

func TestTakeTooMuch(t *testing.T) {
	slice := []int{1, 2, 3}
	s := AsStream(slice)
	s = s.Take(4)
	newSlice, err := AsSlice[int](s)

	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if len(newSlice) != 3 {
		t.Errorf("expected one return value, but got: %v", len(newSlice))
	}
}

func TestTakeNegative(t *testing.T) {
	slice := []int{1, 2, 3}
	s := AsStream(slice)
	s = s.Take(-1)
	_, err := AsSlice[int](s)

	if err == nil {
		t.Errorf("expected to get error, but error was nil")
	}
}

func TestTakeZero(t *testing.T) {
	slice := []int{1, 2, 3}
	s := AsStream(slice)
	s = s.Take(0)
	newSlice, err := AsSlice[int](s)

	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if len(newSlice) != 0 {
		t.Errorf("expected one return value, but got: %v", len(newSlice))
	}
}
