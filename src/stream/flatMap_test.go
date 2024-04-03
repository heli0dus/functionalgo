package stream

import (
	"fmt"
	"slices"
	"testing"
)

func TestFlatMapOk(t *testing.T) {
	slice := []string{"42", "1337", "69"}
	newSlice, err := AsSlice[rune](AsStream(slice).FlatMap(func(v string) []rune { return []rune(v) }))
	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if len(newSlice) != 8 {
		t.Errorf("expected slice of size %v, but got: %v", 8, len(newSlice))
	} else {
		expected := []rune{'4', '2', '1', '3', '3', '7', '6', '9'}
		if !slices.Equal(newSlice, expected) {
			t.Errorf("expected %v, but got: %v", expected, newSlice)
		}
	}
}

func TestFlatMapOkWithErr(t *testing.T) {
	slice := []string{"42", "1337", "69"}
	newSlice, err := AsSlice[rune](AsStream(slice).FlatMap(func(v string) []rune { return []rune(v) }))
	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if len(newSlice) != 8 {
		t.Errorf("expected slice of size %v, but got: %v", 8, len(newSlice))
	} else {
		expected := []rune{'4', '2', '1', '3', '3', '7', '6', '9'}
		if !slices.Equal(newSlice, expected) {
			t.Errorf("expected %v, but got: %v", expected, newSlice)
		}
	}
}

func TestFlatMapWrongNumberOfArgs(t *testing.T) {
	slice := []string{"42", "1337", "69"}
	_, err := AsSlice[rune](AsStream(slice).FlatMap(func(v string, k string) []rune { return []rune(v + k) }))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else {
		t.Log(err)
	}
}

func TestFlatMapWrongTypeOfArg(t *testing.T) {
	slice := []string{"42", "1337", "69"}
	_, err := AsSlice[rune](AsStream(slice).FlatMap(func(v []rune) []rune { return v }))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else {
		t.Log(err)
	}
}

func TestFlatMapWrongNumberOfReturnValues(t *testing.T) {
	slice := []string{"42", "1337", "69"}
	_, err := AsSlice[rune](AsStream(slice).FlatMap(func(v string) ([]rune, int, int) { return []rune(v), 0, 0 }))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else {
		t.Log(err)
	}
}

func TestFlatMapWrongReturnType(t *testing.T) {
	slice := []string{"42", "1337", "69"}
	_, err := AsSlice[rune](AsStream(slice).FlatMap(func(v string) rune { return rune('a') }))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else {
		t.Log(err)
	}
}

func TestFlatMapWrongCast(t *testing.T) {
	slice := []string{"42", "1337", "69"}
	_, err := AsSlice[string](AsStream(slice).FlatMap(func(v string) []rune { return []rune(v) }))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else {
		t.Log(err)
	}
}

func TestFlatMapErr(t *testing.T) {
	slice := []string{"42", "1337", "69"}
	_, err := AsSlice[rune](AsStream(slice).FlatMap(func(v string) ([]rune, error) { return nil, fmt.Errorf("custom error") }))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else if err.Error() != "custom error" {
		t.Errorf("expected error '%v', but got '%v'", "custom error", err.Error())
	}
}
