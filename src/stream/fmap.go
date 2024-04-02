package stream

import "reflect"

func (s Stream) Fmap(f interface{}) Stream {
	// TODO: add param type check
	// TODO: add return type check (val or (val, err))
	res := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(f).Out(0)), 0, s.len)
	for i := 0; i < s.len; i++ {
		args := append([]reflect.Value{}, s.value.Index(i))
		res = reflect.Append(res, reflect.ValueOf(f).Call(args)[0])
	}
	s.slice = res.Interface()
	return s
}
