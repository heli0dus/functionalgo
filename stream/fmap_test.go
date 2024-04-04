package stream

import (
	"fmt"
	"slices"
	"testing"
)

func TestFmapOk(t *testing.T) {
	slice := []int{1, 2, 3}
	newSlice, err := AsSlice[float32](AsStream(slice).Fmap(func(v int) float32 { return float32(v) / 2 }))
	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if len(newSlice) != 3 {
		t.Errorf("expected slice of size %v, but got: %v", 3, len(newSlice))
	} else {
		expected := []float32{0.5, 1, 1.5}
		if !slices.Equal(newSlice, expected) {
			t.Errorf("expected %v, but got: %v", expected, newSlice)
		}
	}
}

func TestFmapOkWithErr(t *testing.T) {
	slice := []int{1, 2, 3}
	newSlice, err := AsSlice[float32](AsStream(slice).Fmap(func(v int) (float32, error) { return float32(v) / 2, nil }))
	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if len(newSlice) != 3 {
		t.Errorf("expected slice of size %v, but got: %v", 3, len(newSlice))
	} else {
		expected := []float32{0.5, 1, 1.5}
		if !slices.Equal(newSlice, expected) {
			t.Errorf("expected %v, but got: %v", expected, newSlice)
		}
	}
}

func TestFmapWrongNumberOfArgs(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := AsSlice[float32](AsStream(slice).Fmap(func(v int, x int) float32 { return float32(v+x) / 2 }))
	if err == nil {
		t.Errorf("expected error, but got none")
	}
}

func TestFmapWrongTypeOfArg(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := AsSlice[float32](AsStream(slice).Fmap(func(v float32) float32 { return float32(v) / 2 }))
	if err == nil {
		t.Errorf("expected error, but got none")
	}
}

func TestFmapWrongNumberOfReturnValues(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := AsSlice[float32](AsStream(slice).Fmap(func(v int) (float32, int, int) { return float32(v) / 2, 1, 2 }))
	if err == nil {
		t.Errorf("expected error, but got none")
	}
}

func TestFmapWrongReturnType(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := AsSlice[float32](AsStream(slice).Fmap(func(v int) (float32, int) { return float32(v) / 2, 1 }))
	if err == nil {
		t.Errorf("expected error, but got none")
	}
}

func TestFmapWrongCast(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := AsSlice[int](AsStream(slice).Fmap(func(v int) float32 { return float32(v) / 2 }))
	if err == nil {
		t.Errorf("expected error, but got none")
	}
}

func TestFmapErr(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := AsSlice[float32](AsStream(slice).Fmap(func(v int) (float32, error) { return 0.0, fmt.Errorf("custom error") }))
	if err == nil {
		t.Errorf("expected error, but got none")
	}
}
