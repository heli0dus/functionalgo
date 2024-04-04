package stream

import (
	"fmt"
	"testing"
)

func TestAnyOkTrue(t *testing.T) {
	slice := []int{1, 2, 3}
	res, err := As[bool](AsStream(slice).Any(func(v int) bool { return v < 2 }))
	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if !res {
		t.Errorf("expected true, but got false")
	}
}

func TestAnyOkFalse(t *testing.T) {
	slice := []int{1, 2, 3}
	res, err := As[bool](AsStream(slice).Any(func(v int) bool { return v < 1 }))
	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if res {
		t.Errorf("expected false, but got true")
	}
}

func TestAnyOkWithErr(t *testing.T) {
	slice := []int{1, 2, 3}
	res, err := As[bool](AsStream(slice).Any(func(v int) (bool, error) { return v < 2, nil }))
	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if !res {
		t.Errorf("expected true, but got false")
	}
}

func TestAnyWrongNumberOfArgs(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := As[bool](AsStream(slice).Any(func(v int, _ int) bool { return v < 4 }))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else {
		t.Log(err.Error())
	}
}

func TestAnyWrongTypeOfArg(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := As[bool](AsStream(slice).Any(func(v float32) bool { return v < 4 }))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else {
		t.Log(err.Error())
	}
}

func TestAnyWrongNumberOfReturnValues(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := As[bool](AsStream(slice).Any(func(v int) (bool, int, error) { return v < 4, 0, nil }))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else {
		t.Log(err.Error())
	}
}

func TestAnyWrongReturnType(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := As[bool](AsStream(slice).Any(func(v int) int { return v }))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else {
		t.Log(err.Error())
	}
}

func TestAnyErr(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := As[bool](AsStream(slice).Any(func(v int) (bool, error) { return false, fmt.Errorf("custom error") }))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else if err.Error() != "custom error" {
		t.Errorf("expercted error '%v', but got %v", "custom error", err.Error())
	}
}

func TestAnyNotFunc(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := As[bool](AsStream(slice).Any(1))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else {
		t.Log(err.Error())
	}
}
