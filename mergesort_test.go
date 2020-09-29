package mysorts

import (
	"testing"
	"sort"
)

func Test_MergeSort(t * testing.T) {
	taco := CreateAndFillSliceWithRandomInts(100)
	MergeSort(&taco, 0, len(taco) -1)

	if !sort.SliceIsSorted(taco, func(p, q int) bool {  
		return taco[p] < taco[q] }) {
		t.Error("Slice with 100 elements not sorted")
	}

	taco = CreateAndFillSliceWithRandomInts(1000)
	MergeSort(&taco, 0, len(taco) -1)

	if !sort.SliceIsSorted(taco, func(p, q int) bool {  
		return taco[p] < taco[q] }) {
		t.Error("Slice with 1000 elements not sorted")
	}

	taco = CreateAndFillSliceWithRandomInts(10000)
	MergeSort(&taco, 0, len(taco) -1)

	if !sort.SliceIsSorted(taco, func(p, q int) bool {  
		return taco[p] < taco[q] }) {
		t.Error("Slice with 10000 elements not sorted")
	}
}