package stream

import (
	"fmt"
	"testing"
)

func TestReduceOk(t *testing.T) {
	slice := []int{1, 2, 3}
	s := AsStream(slice)
	s = s.Fmap(func(v int) float32 { return float32(v) / 2 }).Reduce(func(a int, v float32) int {
		if v > 0.7 {
			return a + 1
		} else {
			return a
		}
	}, 0)
	newSlice, err := AsSlice[int](s)

	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if len(newSlice) != 1 {
		t.Errorf("expected one return value, but got: %v", len(newSlice))
	} else if newSlice[0] != 2 {
		t.Errorf("expected %v, but got: %v", 2, newSlice[0])
	}
}

func TestReduceOkWithErr(t *testing.T) {
	slice := []int{1, 2, 3}
	s := AsStream(slice)
	s = s.Fmap(func(v int) float32 { return float32(v) / 2 }).Reduce(func(a int, v float32) (int, error) {
		if v > 0.7 {
			return a + 1, nil
		} else {
			return a, nil
		}
	}, 0)
	newSlice, err := AsSlice[int](s)

	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if len(newSlice) != 1 {
		t.Errorf("expected one return value, but got: %v", len(newSlice))
	} else if newSlice[0] != 2 {
		t.Errorf("expected %v, but got: %v", 2, newSlice[0])
	}
}

func TestReduceErr(t *testing.T) {
	slice := []int{1, 2, 3}
	s := AsStream(slice)
	s = s.Fmap(func(v int) float32 { return float32(v) / 2 }).Reduce(func(a int, v float32) (int, error) {
		if v > 0.7 {
			return 0, fmt.Errorf("custom error")
		} else {
			return a, nil
		}
	}, 0)
	_, err := AsSlice[int](s)

	if err == nil {
		t.Error("expected error, but got none")
	} else if err.Error() != "custom error" {
		t.Errorf("expected error message: '%v', but got '%v'", "custom error", err.Error())
	}
}

func TestReduceWrongInitType(t *testing.T) {
	slice := []int{1, 2, 3}
	s := AsStream(slice)
	s = s.Fmap(func(v int) float32 { return float32(v) / 2 }).Reduce(func(a int, v float32) int {
		if v > 0.7 {
			return a + 1
		} else {
			return a
		}
	}, "init")
	_, err := AsSlice[int](s)
	if err == nil {
		t.Errorf("expected error, but got none")
	}
}
