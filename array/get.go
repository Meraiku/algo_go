package array

import (
	"math/rand"
)

const Size = 50000

func Get() []int {
	arr := make([]int, Size)

	for i := range arr {
		arr[i] = rand.Intn(Size)
	}
	return arr
}
