package stream

import (
	"fmt"
	"reflect"
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
	res := reflect.MakeSlice(reflect.SliceOf(s.elemType), 0, resSize)

	for i := n; i < s.len; i++ {
		res = reflect.Append(res, s.value.Index(i))
	}
	s.len = resSize
	s.value = res
	s.slice = res.Interface()
	return s
}
