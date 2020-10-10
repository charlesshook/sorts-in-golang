package mysorts

func InsertionSort(myslice *[]int) {
	for i := 0; i < len(*myslice); i++ {
		key := (*myslice)[i]
		j := i

		for j > 0 && (*myslice)[j-1] > key {
			(*myslice)[j] = (*myslice)[j-1]
			j--
		}

		(*myslice)[j] = key
	}
}
