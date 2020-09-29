package mysorts

func MergeSort(myslice *[]int, p int, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(myslice, p, q)
		MergeSort(myslice, q +1, r)
		Merge(myslice, p, q, r)
	}
}

func Merge(myslice *[]int, p int, q int, r int) {
	n1 := q - p +1
	n2 := r - q
	var L, R []int

	for i := 1; i < n1; i++ {
		L = append(L, (*myslice)[p + i -1])
	}

	for j := 1; j < n2; j++ {
		R = append(R, (*myslice)[q + j])
	}

	L = append(L, int(^uint(0) >> 1))
	R = append(R, int(^uint(0) >> 1))

	i, j := 1, 1

	for k := p; k < r; k++ {
		if L[i] <= R[j] {
			(*myslice)[k] = L[i]
			i++
		} else if (*myslice)[k] == R[j] {
			j++
		}
	}
}