package main

import (
	"fmt"
	"time"

	use "use/array"
)

// UseArray

func main() {
	length := 10_000_000
	myArray := make([]int, 0, 10)
	// make rundom values
	for i := 0; i < length; i++ {
		item := int(time.Now().UnixNano() / 1000000)
		myArray = append(myArray, item)
	}

	startTime := time.Now()

	// myArray := []int{10, 2, 30, 4, 5, 6, 7, 8, 9}

	// fmt.Println("Start: ", myArray)
	newArray := use.NewArray(myArray).
		Filter(func(item any, index int) bool {
			return item.(int) >= 10
		}).
		Map(func(item any, index int) any {
			return item.(int) + 1
		}).
		Reduce(
			func(item any, i int, acc any) any {
				return acc.(int) + item.(int)
			},
			0,
		).
		Run()

	endTime := time.Now()

	fmt.Println("Time: ", endTime.Sub(startTime))
	fmt.Printf("Type: %T\n", newArray)
	fmt.Println("Result: ", newArray)

	// err := fmt.Errorf("oh noes: %v", os.ErrNotExist)
	// fmt.Errorf("oh noes: %v", err)
}
