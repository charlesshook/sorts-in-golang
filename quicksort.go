package mysorts

func QuickSort(myslice *[]int, left int, right int) {
	index := partition(myslice, left, right)

	if left < index - 1 {
		QuickSort(myslice, left, index -1)
	}

	if index < right {
		QuickSort(myslice, index, right)
	}
}

func partition(myslice *[]int, left int, right int) int {
	pivot := (*myslice)[left + (right - left) / 2]

	for left <= right {
		for (*myslice)[left] < pivot {
			left++
		}

		for (*myslice)[right] > pivot {
			right--
		}

		if left <= right {
			temp := (*myslice)[left]
			(*myslice)[left] = (*myslice)[right]
			(*myslice)[right] = temp

			left++
			right--
		}
	}

	return left
}