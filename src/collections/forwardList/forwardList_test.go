package forwardList

import "testing"

func TestFmap(t *testing.T) {
	initialSlice := []int{1, 2, 3}
	list := fromSlice(initialSlice)

	newList := Fmap(list, func(x int) int {
		return x + 1
	})

	expectedSlice := []int{2, 3, 4}
	resultSlice := newList.toSlice()

	for i, val := range resultSlice {
		if val != expectedSlice[i] {
			t.Errorf("Expected slice %v, but got %v", expectedSlice, resultSlice)
			break
		}
	}
}

func TestReduce(t *testing.T) {
	initialSlice := []int{1, 2, 3}
	list := fromSlice(initialSlice)

	result := Reduce(list, 0, func(accumulator int, value int) int {
		return accumulator + value
	})

	expectedSum := 6
	if result != expectedSum {
		t.Errorf("Expected sum %d, but got %d", expectedSum, result)
	}
}

func TestFromSlice(t *testing.T) {
	initialSlice := []int{1, 2, 3}
	list := fromSlice(initialSlice)
	resultSlice := list.toSlice()

	for i, val := range resultSlice {
		if val != initialSlice[i] {
			t.Errorf("Expected slice %v, but got %v", initialSlice, resultSlice)
			break
		}
	}
}

func TestToSlice(t *testing.T) {
	initialSlice := []int{1, 2, 3}
	list := fromSlice(initialSlice)

	resultSlice := list.toSlice()
	for i, val := range resultSlice {
		if val != initialSlice[i] {
			t.Errorf("Expected slice %v, but got %v", initialSlice, resultSlice)
			break
		}
	}
}
