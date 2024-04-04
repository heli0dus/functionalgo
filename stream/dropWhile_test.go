package stream

import "testing"

func TestDropWhileUnchanged(t *testing.T) {
	arr := []int{3, 2, 4, 1, 6}
	s := AsStream(arr)
	s = s.DropWhile(func(x int) bool { return false })
	res, err := AsSlice[int](s)

	if err != nil {
		t.Errorf("error were not expected, but got %v", err.Error())
	} else if len(res) != len(arr) {
		t.Errorf("Got '%v' while expected '%v'", len(res), len(arr))
	}
}

func TestDropWhileConstraint(t *testing.T) {
	arr := []int{3, 2, 4, 1, 6}
	s := AsStream(arr)
	s = s.DropWhile(func(x int) bool { return true })
	res, err := AsSlice[int](s)

	if err != nil {
		t.Errorf("error were not expected, but got %v", err.Error())
	} else if len(res) != 0 {
		t.Errorf("Got '%v' while expected '%v'", len(res), 0)
	}
}

func TestDropWhileRegular(t *testing.T) {
	arr := []int{3, 2, 4, 1, 6}
	result := []int{2, 4, 1, 6}

	s := AsStream(arr)
	s = s.DropWhile(func(x int) bool { return x%2 != 0 })
	res, err := AsSlice[int](s)
	if err != nil {
		t.Errorf("error were not expected, but got %v", err.Error())
	} else if len(res) != len(result) {
		t.Errorf("Got '%v' while expected '%v'", len(res), len(result))
		return
	}
	for i := 0; i < len(res); i++ {
		if res[i] != result[i] {
			t.Errorf("Got '%v' while expected '%v'", res[i], result[i])
		}
	}
}
