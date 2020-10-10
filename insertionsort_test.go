package mysorts

import (
	"sort"
	"testing"
)

func Test_InsertionSort(t *testing.T) {
	taco := CreateAndFillSliceWithRandomInts(100)
	InsertionSort(&taco)

	if !sort.SliceIsSorted(taco, func(p, q int) bool {
		return taco[p] < taco[q]
	}) {
		t.Error("Slice with 100 elements not sorted")
	}

	taco = CreateAndFillSliceWithRandomInts(1000)
	InsertionSort(&taco)

	if !sort.SliceIsSorted(taco, func(p, q int) bool {
		return taco[p] < taco[q]
	}) {
		t.Error("Slice with 1000 elements not sorted")
	}

	taco = CreateAndFillSliceWithRandomInts(10000)
	InsertionSort(&taco)

	if !sort.SliceIsSorted(taco, func(p, q int) bool {
		return taco[p] < taco[q]
	}) {
		t.Error("Slice with 10000 elements not sorted")
	}
}
