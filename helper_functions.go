package mysorts

import (
	"math/rand"
	"time"
)

func CreateAndFillSliceWithRandomInts(numberOfElements int) []int {
	var myslice []int
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numberOfElements; i++ {
		myslice = append(myslice, rand.Int())
	}

	return myslice
}