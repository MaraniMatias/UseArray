package array

import "testing"

func TestArrayOperations(t *testing.T) {
	myArray := []int{10, 2, 30, 4, 5, 6, 7, 8, 9}
	expectedResult := 92 // (10+1) + (30+1) + (4+1) + (5+1) + (6+1) + (7+1) + (8+1) + (9+1)

	reducedResult := NewArray(myArray).Filter(func(item any, index int) bool {
		return item.(int) >= 10
	}).Map(func(item any, index int) any {
		return item.(int) + 1
	}).Reduce(func(item any, i int, acc any) any {
		return acc.(int) + item.(int)
	}, 0).Run()

	if reducedResult != expectedResult {
		t.Errorf("Expected %d but got %d", expectedResult, reducedResult)
	}
}

func BenchmarkArray(b *testing.B) {
	myArray := []int{10, 2, 30, 4, 5, 6, 7, 8, 9, 23, 134, 23}

	for n := 0; n < b.N; n++ {
		_ = NewArray(myArray).
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
	}
}
