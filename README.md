# UseArray

This package is a wrapper for the [array](https://golang.org/pkg/sort/#Slice) package.

## Provide

```go
type (
 // FilterFn is a function for filter
 Filter func(item any, index int) bool
 // MapFn is a function for map
 Map func(item any, index int) any
 // ReduceFn is a function for reduce
 Reduce func(item any, i int, acc any) any
 // ForEachFn is a function for forEach
 ForEach[T any] func(item T, i int, arr []T)
 // EveryFn is a function for every
 Every[T any] func(item T, i int) bool
 // SomeFn is a function for some, return true if some item is true
 Some[T any] func(item T, i int) bool
 // FindFn is a function for find, return ele, index
 Find[T any] func(item T, i int) (T, int)
 // FindLastFn is a function for findLast,  return ele, index
 FindLast[T any] func(item T, i int) (T, int)
 // SortFn is a function for sort, return 1 if i > j
 Sort[T any] func(arr []T, i int, j int) uint8
```

## Example

```go

package main

import (
 "fmt"

 use "use/array"
)

func main() {
 myArray := []int{10, 2, 30, 4, 5, 6, 7, 8, 9}
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

 fmt.Println("Result: ", newArray) // 92
}
```
