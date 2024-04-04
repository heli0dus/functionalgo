package stream

import (
	"slices"
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
		t.Errorf("Got '%v' while expected '%v'", len(newSlice), 2)
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
		t.Errorf("Got '%v' while expected '%v'", len(newSlice), 3)
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
		t.Errorf("Got '%v' while expected '%v'", newSlice, 0)
	} else if !slices.Equal([]int{}, newSlice) {
		t.Errorf("Got '%v' while expected '%v'", newSlice, slice)
	}
}

func TestTakeByOne(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := AsStream(arr)
	for i := len(arr); i > 0; i-- {
		s = s.Take(i)
		res, err := AsSlice[int](s)

		if err != nil {
			t.Errorf("error were not expected, but got %v", err.Error())
		} else if !slices.Equal(arr[:i], res) {
			t.Errorf("Got '%v' while expected '%v'", res, arr[i:])
		}
	}

}
