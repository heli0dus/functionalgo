package stream

import (
	"fmt"
	"reflect"
)

func (s Stream) Take(n int) Stream {
	// TODO: replace with s.Valid()
	if s.err != nil {
		return s
	}
	if n < 0 {
		return s.Error(fmt.Errorf("can't take negative number of elements: %d", n))
	}

	resSize := min(s.len, n)
	res := reflect.MakeSlice(reflect.SliceOf(s.elemType), 0, resSize)

	for i := 0; i < resSize; i++ {
		res = reflect.Append(res, s.value.Index(i))
	}

	s.len = resSize
	s.value = res
	s.slice = res.Interface()
	return s
}
