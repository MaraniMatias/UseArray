package array

// type (
// 	FilterFn[T any]        func(item T, index int) bool
// 	MapFn[T any]           func(item T, index int) T
// 	ReduceFn[T any, P any] func(item T, i int, acc P) P
// )

type (
	// FilterFn is a function for filter
	FilterFn func(item any, index int) bool
	// MapFn is a function for map
	MapFn func(item any, index int) any
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
	fn     interface{}
	args   []any
}

type array[T any] struct {
	arr         []T
	listOfFuncs []typeFuc
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
	args := []any{init}
	a.listOfFuncs = append(a.listOfFuncs, typeFuc{"reduce", fn, args})
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

// func reduce[T any, P any](arr []T, fn func(item T, index int, acc P) P, initAcc P) P {
func reduce[T any](arr []T, fn func(item any, index int, acc any) any, initAcc any) any {
	length := len(arr)
	acc := initAcc
	for i := 0; i < length; i++ {
		acc = fn(arr[i], i, acc)
	}
	return acc
}

func (a *array[T]) Run() interface{} {
	length := len(a.arr)
	state := make([]T, 0, length)

	var initAcc any
	var reduceFn ReduceFn
	last := a.listOfFuncs[len(a.listOfFuncs)-1]
	if last.typeFn == "reduce" {
		reduceFn = last.fn.(ReduceFn)
		initAcc = last.args[0]
	} else {
		reduceFn = func(item any, i int, acc any) any {
			return append(acc.([]any), item)
		}
	}

	for i := 0; i < length; i++ {
		item := a.arr[i]
		result := every[typeFuc](a.listOfFuncs, func(tFn any, index int) bool {
			tfn := tFn.(typeFuc)
			if tfn.typeFn == "map" {
				fn := tfn.fn.(MapFn)
				item = fn(item, index).(T)
				return true
			} else if tfn.typeFn == "filter" {
				fn := tfn.fn.(FilterFn)
				return fn(item, index)
			} else if tfn.typeFn == "reduce" {
				reduceFn = tfn.fn.(ReduceFn)
				initAcc = tfn.args[0]
				return true
			}
			return false
		})
		if result {
			state = append(state, item)
		}
	}

	if initAcc == nil {
		initAcc = make([]T, 0, len(state))
	}

	// func reduce[T any, P any](arr []T, fn func(item T, index int, acc P) P, initAcc P) P {
	// return reduce[T, []T](state.([]T), reduceFn.(ReduceFn), initAcc.([]T)).([]T)
	// return reduce(state, reduceFn, initAcc).([]T)
	return reduce(state, reduceFn, initAcc).(T)
}

// UseArray is a constructor for Array
func UseArray[T any](arr []T) Array {
	a := &array[T]{arr: arr}
	return a
}
