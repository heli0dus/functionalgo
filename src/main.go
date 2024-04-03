package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/heli0dus/functionalgo/src/stream"
)

func isEven(i int) bool {
	return i%2 == 0
}

func square(i int) int {
	return i * i
}

func add(a int, b int) int {
	return a + b
}

func main() {
	coolNumbers := []string{
		"42",
		"69",
		"13",
		"666",
		"1337",
		"1",
		"0",
		"-5",
		"12",
	}
	s := stream.AsStream(coolNumbers)
	s = s.Fmap(strconv.Atoi).Filter(isEven).Fmap(square).Reduce(add, 0)
	res, err := stream.As[int](s)
	if err != nil {
		fmt.Printf("error: %v", err.Error())
		os.Exit(1)
	}
	// correct answer is 445464
	fmt.Printf("sum of squares of even numbers: %v\n", res)
}
