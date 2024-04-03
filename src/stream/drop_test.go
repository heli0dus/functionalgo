package stream

import "testing"

func TestDrop(t *testing.T) {
	arr := []int{1, 2, 3}
	s := AsStream(arr)
	s = s.Drop(2)
	res, err := AsSlice[int](s)

	if err != nil {
		t.Errorf("error were not expected, but got %v", err.Error())
	} else if len(res) != 1 {
		t.Errorf("Got '%v' while expected '%v'", len(res), 1)
	} else if res[0] != 3 {
		t.Errorf("Got '%v' while expected '%v'", res[0], 3)
	}
}

func TestDropTooMuch(t *testing.T) {
	arr := []int{1, 2, 3}
	s := AsStream(arr)
	s = s.Drop(4)
	res, err := AsSlice[int](s)

	if err != nil {
		t.Errorf("error were not expected, but got %v", err.Error())
	} else if len(res) != 0 {
		t.Errorf("Got '%v' while expected '%v'", res[0], 0)
	}
}

func TestDropNegative(t *testing.T) {
	arr := []int{1, 2, 3}
	s := AsStream(arr)
	s = s.Drop(-1)
	_, err := AsSlice[int](s)

	if err == nil {
		t.Errorf("error expected, but got %v", nil)

	}
}

func TestDropZero(t *testing.T) {
	arr := []int{1, 2, 3}
	s := AsStream(arr)
	s = s.Drop(0)
	res, err := AsSlice[int](s)

	if err != nil {
		t.Errorf("error were not expected, but got %v", err.Error())
	} else if len(res) != len(arr) {
		t.Errorf("Got '%v' while expected '%v'", res[0], 0)
	}

	for i := 0; i < len(res); i++ {
		if res[i] != arr[i] {
			t.Errorf("Got '%v' while expected '%v'", res[i], arr[i])
		}
	}
}
