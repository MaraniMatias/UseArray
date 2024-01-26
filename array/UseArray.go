package array

type (
	// FilterFn is a function for filter
	FilterFn func(item any, index int) bool
	// FilterFn[T any] func(item T, index int) bool
	// MapFn is a function for map
	MapFn func(item any, index int) any
	// MapFn[T any] func(item T, index int) T
	// ReduceFn is a function for reduce
	ReduceFn func(item any, i int, acc any) any
)

// Array is a wrapper for go slices
type Array interface {
	Filter(cb FilterFn) Array
	Map(cb MapFn) Array
	Reduce(cb ReduceFn, init any) Array
	// ReduceRight(func(arr []int, i int, acc int) int, initAcc int) *Array
	// Sort(func(arr []int, i int, j int) bool) *Array
	// ForEach(func(arr []int, i int)) *Array
	// Find(func(arr []int, i int) bool) *Array
	// FindIndex(func(arr []int, i int) bool) *Array
	// FindLast(func(arr []int, i int) bool) *Array
	// FindLastIndex(func(arr []int, i int) bool) *Array
	// Every(func(arr []int, i int) bool) *Array
	// Some(func(arr []int, i int) bool) *Array
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

func (a *array[T]) Filter(fn FilterFn) Array {
	a.listOfFuncs = append(a.listOfFuncs, typeFuc{"filter", fn, nil})
	return a
}

func (a *array[T]) Map(fn MapFn) Array {
	a.listOfFuncs = append(a.listOfFuncs, typeFuc{"map", fn, nil})
	return a
}

func (a *array[T]) Reduce(fn ReduceFn, init any) Array {
	a.reduceFn.fn = fn
	a.reduceFn.init = init
	return a
}

// func reduce[T any, P any](arr []T, fn func(item T, index int, acc P) P, initAcc P) P {
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
			return append(acc.([]any), item)
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

// UseArray is a constructor for Array
func UseArray[T any](arr []T) Array {
	a := &array[T]{arr: arr}
	return a
}
