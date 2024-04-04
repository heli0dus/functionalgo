package stream

import (
	"fmt"
)

func (s Stream) Take(n int) Stream {
	if !s.Valid() {
		return s
	}
	if n < 0 {
		return s.Error(fmt.Errorf("can't take negative number of elements: %d", n))
	}

	resSize := min(s.len, n)

	s.value = s.value.Slice(0, resSize)
	s.slice = s.value.Interface()

	s.len = resSize

	return s
}
