package stream

import (
	"fmt"
)

func (s Stream) Drop(n int) Stream {
	if !s.Valid() {
		return s
	}

	if n < 0 {
		return s.Error(fmt.Errorf("can't call `Drop` with negative number of elements: %d", n))
	}

	n = min(s.len, n)
	resSize := s.len - n

	s.value = s.value.Slice(n, s.len)
	s.slice = s.value.Interface()

	s.len = resSize

	return s
}
