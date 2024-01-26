package array

// type (
// 	FilterFn[T any]        func(item T, index int) bool
// 	MapFn[T any]           func(item T, index int) T
// 	ReduceFn[T any, P any] func(arr []T, i int, acc P) T
// )

type (
	FilterFn func(item any, index int) bool
	MapFn    func(item any, index int) any
	ReduceFn func(arr []any, i int, acc any) any
)

// Array is a wrapper for go slices
type Array interface {
	Filter(cb FilterFn) Array
	Map(cb MapFn) Array
	Reduce(cb ReduceFn) Array
	// ReduceRight(func(arr []int, i int, acc int) int, initAcc int) *Array
	// Sort(func(arr []int, i int, j int) bool) *Array
	// ForEach(func(arr []int, i int)) *Array
	// Find(func(arr []int, i int) bool) *Array
	// FindIndex(func(arr []int, i int) bool) *Array
	// FindLast(func(arr []int, i int) bool) *Array
	// FindLastIndex(func(arr []int, i int) bool) *Array
	// Every(func(arr []int, i int) bool) *Array
	// Some(func(arr []int, i int) bool) *Array
	Run() (interface{}, error)
}

type typeFuc struct {
	typeFn string
	fn     interface{}
}

type array[T any] struct {
	arr         []T
	listOfFuncs []typeFuc
}

func (a *array[T]) Filter(fn FilterFn) Array {
	a.listOfFuncs = append(a.listOfFuncs, typeFuc{"filter", fn})
	return a
}

func (a *array[T]) Map(fn MapFn) Array {
	a.listOfFuncs = append(a.listOfFuncs, typeFuc{"map", fn})
	return a
}

func (a *array[T]) Reduce(fn ReduceFn) Array {
	a.listOfFuncs = append(a.listOfFuncs, typeFuc{"reduce", fn})
	return a
}

func every[T any](arr []T, fn FilterFn) bool {
	length := len(arr)
	for i := 0; i < length; i++ {
		if !fn(arr[i], i) {
			return false
		}
	}
	return true
}

func reduce[T any](arr []any, fn ReduceFn, initAcc T) T {
	length := len(arr)
	acc := initAcc
	for i := 0; i < length; i++ {
		acc = fn(arr, i, acc).(T)
	}
	return acc
}

func (a *array[T]) Run() (interface{}, error) {
	length := len(a.arr)
	count := 0
	state := make([]T, 0, length)

	for i := 0; i < length; i++ {
		item := a.arr[i]
		result := every[typeFuc](a.listOfFuncs, func(tFn any, index int) bool {
			tfn := tFn.(typeFuc)
			if tfn.typeFn == "map" {
				fn := tfn.fn.(MapFn)
				item = fn(item, index).(T)
				// state = append(state, fn(item, i).(int))
				// count++
				return true
			} else if tfn.typeFn == "filter" {
				fn := tfn.fn.(FilterFn)
				// state = append(state, item)
				// count++
				return fn(item, index)
			}
			return false
		})
		if result {
			state = append(state, item)
			count++
		}
	}
	return state, nil
}

// UseArray is a constructor for Array
func UseArray[T any](arr []T) Array {
	a := &array[T]{arr: arr}
	return a
}
