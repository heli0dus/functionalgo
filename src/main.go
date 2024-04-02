package main

import (
	"fmt"

	"github.com/heli0dus/functionalgo/src/stream"
)

func main() {
	slice := []int{1, 2, 3}
	newSlice := stream.AsSlice[float32](stream.AsStream(slice).Fmap(func(v int) float32 { return float32(v) / 2 }))
	fmt.Println(len(newSlice))
	for i := 0; i < len(newSlice); i++ {
		fmt.Println(newSlice[i])
	}
}
