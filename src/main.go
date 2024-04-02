package main

import (
	"fmt"

	"github.com/heli0dus/functionalgo/src/stream"
)

func testFmapOk() {
	slice := []int{1, 2, 3}
	newSlice, err := stream.AsSlice[float32](stream.AsStream(slice).Fmap(func(v int) float32 { return float32(v) / 2 }))
	if err != nil {
		fmt.Println(err.Error())
		panic("testFmapOk failed")
	} else {
		fmt.Println(len(newSlice))
		for i := 0; i < len(newSlice); i++ {
			fmt.Print(newSlice[i], " ")
		}
		fmt.Println()
	}
}

func testFmapOkWithErr() {
	slice := []int{1, 2, 3}
	newSlice, err := stream.AsSlice[float32](stream.AsStream(slice).Fmap(func(v int) (float32, error) { return float32(v) / 2, nil }))
	if err != nil {
		fmt.Println(err.Error())
		panic("testFmapWrongCast failed")
	} else {
		fmt.Println(len(newSlice))
		for i := 0; i < len(newSlice); i++ {
			fmt.Print(newSlice[i], " ")
		}
		fmt.Println()
	}
}

func testFmapWrongNumberOfArgs() {
	slice := []int{1, 2, 3}
	_, err := stream.AsSlice[float32](stream.AsStream(slice).Fmap(func(v int, x int) float32 { return float32(v+x) / 2 }))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		panic("testFmapWrongNumberOfArgs failed")
	}
}

func testFmapWrongTypeOfArg() {
	slice := []int{1, 2, 3}
	_, err := stream.AsSlice[float32](stream.AsStream(slice).Fmap(func(v float32) float32 { return float32(v) / 2 }))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		panic("testFmapWrongTypeOfArg failed")
	}
}

func testFmapWrongNumberOfReturnValues() {
	slice := []int{1, 2, 3}
	_, err := stream.AsSlice[float32](stream.AsStream(slice).Fmap(func(v int) (float32, int, int) { return float32(v) / 2, 1, 2 }))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		panic("testFmapWrongNumberOfReturnValues failed")
	}
}

func testFmapWrongReturnType() {
	slice := []int{1, 2, 3}
	_, err := stream.AsSlice[float32](stream.AsStream(slice).Fmap(func(v int) (float32, int) { return float32(v) / 2, 1 }))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		panic("testFmapWrongReturnType failed")
	}
}

func testFmapWrongCast() {
	slice := []int{1, 2, 3}
	_, err := stream.AsSlice[int](stream.AsStream(slice).Fmap(func(v int) float32 { return float32(v) / 2 }))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		panic("testFmapWrongCast failed")
	}
}

func testFmapErr() {
	slice := []int{1, 2, 3}
	_, err := stream.AsSlice[float32](stream.AsStream(slice).Fmap(func(v int) (float32, error) { return 0.0, fmt.Errorf("custom error") }))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		panic("testFmapErr failed")
	}
}

func main() {
	testFmapOk()
	testFmapOkWithErr()
	testFmapWrongNumberOfArgs()
	testFmapWrongTypeOfArg()
	testFmapWrongNumberOfReturnValues()
	testFmapWrongReturnType()
	testFmapWrongCast()
	testFmapErr()
}
