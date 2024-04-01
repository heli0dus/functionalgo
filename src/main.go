package main

import "fmt"

type ArrayList[T any, V any] []T

func (arr ArrayList[T, V]) fmap(f func(T) V) []V {
	res := make([]V, 0, len(arr))
	for _, val := range arr {
		res = append(res, f(val))
	}

	return res
}

func id[T any](x T) T {
	return x
}

func toString(x int) string {
	return fmt.Sprintf("%d", x)
}

func main() {
	a := make([]int, 10)
	var arr ArrayList[int, string] = a
	brr := (ArrayList[int, string])(a).fmap(toString)
	var _ []string = brr
}
