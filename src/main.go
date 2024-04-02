package main

import (
	"fmt"
	"os"

	"github.com/heli0dus/functionalgo/src/stream"
)

func main() {
	slice := []int{1, 2, 3}
	newSlice, err := stream.AsSlice[int](stream.AsStream(slice).Fmap(func(v int) float32 { return float32(v) / 2 }))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(len(newSlice))
	for i := 0; i < len(newSlice); i++ {
		fmt.Println(newSlice[i])
	}
}
