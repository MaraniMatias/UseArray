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

	newArry, err := use.UseArray(myArray).
		Filter(func(item any, index int) bool {
			return item.(int) > 10
		}).
		Map(func(item any, index int) any {
			return item.(int) + 1
		}).
		// Reduce(func(arr []int, i int, acc int) int {
		// 	return acc + arr[i]
		// }, 0).
		Run()
	if err != nil {
		fmt.Println(err)
	}

	endTime := time.Now()

	fmt.Println("Time: ", endTime.Sub(startTime))
	fmt.Println("Result: ", len(newArry.([]int)))
	// for _, v := range newArry.([]int) {
	// 	fmt.Println(v)
	// }

	// err := fmt.Errorf("oh noes: %v", os.ErrNotExist)
	// fmt.Errorf("oh noes: %v", err)
}
