package main

import (
	"fmt"
	"time"

	use "use/array"
)

// UseArray

func main() {
	length := 1000000
	myArray := make([]int, 0, 10)
	// make rundom values
	for i := 0; i < length; i++ {
		item := int(time.Now().UnixNano() / 1000000)
		myArray = append(myArray, item)
	}

	startTime := time.Now()

	// myArray := []int{1, 2, 3, 10, 12, 13}
	newArray := use.UseArray(myArray).
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
	fmt.Println("Result: ", newArray)
	// err := fmt.Errorf("oh noes: %v", os.ErrNotExist)
	// fmt.Errorf("oh noes: %v", err)
}
