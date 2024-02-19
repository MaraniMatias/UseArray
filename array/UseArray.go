package array

type (
	// FilterFn is a function for filter
	FilterFn func(item any, index int) bool
	// MapFn is a function for map
	MapFn func(item any, index int) any
	// ReduceFn is a function for reduce
	ReduceFn func(item any, i int, acc any) any
	// ForEachFn is a function for forEach
	ForEachFn[T any] func(item T, i int, arr []T)
	// EveryFn is a function for every
	EveryFn[T any] func(item T, i int) bool
	// Some is a function for some, return true if some item is true
	SomeFn[T any] func(item T, i int) bool
	// Find is a function for find, return ele, indexñ
	FindFn[T any] func(item T, i int) (T, int)
	// FindLast is a function for findLast,  return ele, index
	FindLastFn[T any] func(item T, i int) (T, int)
	// SortFn is a function for sort, return 1 if i > j
	SortFn[T any] func(arr []T, i int, j int) uint8
)

// A delegate type for filtering data
type ArrayForReduce interface {
	Run() interface{}
}

// Array is a wrapper for go slices
type Array interface {
	Filter(cb FilterFn) Array
	Map(cb MapFn) Array
	Reduce(cb ReduceFn, init any) ArrayForReduce
	// ReduceRight(func(fn ReduceFn, initAcc int)) ArrayForReduce
	// Sort(func(arr []int, i int, j int) bool) Array
	Run() interface{}
}

type typeFuc struct {
	typeFn string
	fn     any
	args   []any
}

type array[T any] struct {
	arr         []T
	listOfFuncs []typeFuc
	reduceFn    struct {
		fn   ReduceFn
		init any
	}
}

func (a *array[T]) ForEach(fn ForEachFn[T]) *array[T] {
	length := len(a.arr)
	for i := 0; i < length; i++ {
		fn(a.arr[i], i, a.arr)
	}
	return a
}

func (a *array[T]) Filter(fn FilterFn) Array {
	a.listOfFuncs = append(a.listOfFuncs, typeFuc{"filter", fn, nil})
	return a
}

func (a *array[T]) Map(fn MapFn) Array {
	a.listOfFuncs = append(a.listOfFuncs, typeFuc{"map", fn, nil})
	return a
}

func (a *array[T]) Reduce(fn ReduceFn, init any) ArrayForReduce {
	a.reduceFn.fn = fn
	a.reduceFn.init = init
	return a
}

func (a *array[T]) Every(fn EveryFn[T]) bool {
	length := len(a.arr)
	for i := 0; i < length; i++ {
		if !fn(a.arr[i], i) {
			return false
		}
	}
	return true
}

func (a *array[T]) Some(fn SomeFn[T]) bool {
	length := len(a.arr)
	for i := 0; i < length; i++ {
		if fn(a.arr[i], i) {
			return true
		}
	}
	return false
}

func (a *array[T]) Find(fn FindFn[T]) (*T, int) {
	length := len(a.arr)
	for i := 0; i < length; i++ {
		if item, index := fn(a.arr[i], i); index != -1 {
			return &item, index
		}
	}
	return nil, -1
}

func (a *array[T]) FindLast(fn FindLastFn[T]) (*T, int) {
	length := len(a.arr)
	for i := length - 1; i >= 0; i-- {
		if item, index := fn(a.arr[i], i); index != -1 {
			return &item, index
		}
	}
	return nil, -1
}

func (a *array[T]) Sort(fn func(arr []T, i int, j int) uint8) *array[T] {
	length := len(a.arr)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if fn(a.arr, i, j) == 1 {
				a.arr[i], a.arr[j] = a.arr[j], a.arr[i]
			}
		}
	}
	return a
}

func reduce[T any](arr []T, fn func(item any, index int, acc any) any, initAcc any) any {
	length := len(arr)
	acc := initAcc
	for i := 0; i < length; i++ {
		acc = fn(arr[i], i, acc)
	}
	return acc
}

func (a *array[T]) Run() any {
	if a.reduceFn.fn == nil {
		a.reduceFn.init = make([]any, 0, len(a.arr))
		a.reduceFn.fn = func(item any, i int, acc any) any {
			// return append(acc.([]any), item)
			acc.([]any)[i] = item
			return acc
		}
	}

	rfn := a.reduceFn.fn
	init := a.reduceFn.init

	return reduce[T](a.arr, func(item any, i int, acc any) any {
		for _, tFn := range a.listOfFuncs {
			t := tFn.typeFn
			fn := tFn.fn

			if t == "filter" && !fn.(FilterFn)(item, i) {
				return acc
			}
			if t == "map" {
				item = fn.(MapFn)(item, i)
			}
		}
		return rfn(item, i, acc)
	}, init)
}

// NewArray is a constructor for Array
func NewArray[T any](arr []T) Array {
	a := &array[T]{arr: arr}
	return a
}
