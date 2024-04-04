package stream

import (
	"fmt"
	"testing"
)

func TestAllOkTrue(t *testing.T) {
	slice := []int{1, 2, 3}
	res, err := As[bool](AsStream(slice).All(func(v int) bool { return v < 4 }))
	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if !res {
		t.Errorf("expected true, but got false")
	}
}

func TestAllOkFalse(t *testing.T) {
	slice := []int{1, 2, 3}
	res, err := As[bool](AsStream(slice).All(func(v int) bool { return v < 3 }))
	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if res {
		t.Errorf("expected false, but got true")
	}
}

func TestAllOkWithErr(t *testing.T) {
	slice := []int{1, 2, 3}
	res, err := As[bool](AsStream(slice).All(func(v int) (bool, error) { return v < 4, nil }))
	if err != nil {
		t.Errorf("expected no error, but got: %v", err.Error())
	} else if !res {
		t.Errorf("expected true, but got false")
	}
}

func TestAllWrongNumberOfArgs(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := As[bool](AsStream(slice).All(func(v int, _ int) bool { return v < 4 }))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else {
		t.Log(err.Error())
	}
}

func TestAllWrongTypeOfArg(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := As[bool](AsStream(slice).All(func(v float32) bool { return v < 4 }))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else {
		t.Log(err.Error())
	}
}

func TestAllWrongNumberOfReturnValues(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := As[bool](AsStream(slice).All(func(v int) (bool, int, error) { return v < 4, 0, nil }))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else {
		t.Log(err.Error())
	}
}

func TestAllWrongReturnType(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := As[bool](AsStream(slice).All(func(v int) int { return v }))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else {
		t.Log(err.Error())
	}
}

func TestAllErr(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := As[bool](AsStream(slice).All(func(v int) (bool, error) { return false, fmt.Errorf("custom error") }))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else if err.Error() != "custom error" {
		t.Errorf("expercted error '%v', but got %v", "custom error", err.Error())
	}
}

func TestAllNotFunc(t *testing.T) {
	slice := []int{1, 2, 3}
	_, err := As[bool](AsStream(slice).All(1))
	if err == nil {
		t.Errorf("expected error, but got none")
	} else {
		t.Log(err.Error())
	}
}
