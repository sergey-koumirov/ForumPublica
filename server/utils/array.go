package utils

import (
	"github.com/twotwotwo/sorts/sortutil"
)

// FindInt64 returns the smallest index i at which x == a[i], or len(a) if there is no such index.
func FindInt64(a []int64, x int64) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

// FindInt32 returns the smallest index i at which x == a[i], or len(a) if there is no such index.
func FindInt32(a []int32, x int32) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

//DiffInt64 diff int64 arrays
func DiffInt64(a, b []int64) []int64 {
	a = sortIfNeeded(a)
	b = sortIfNeeded(b)
	var d []int64
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
		} else if a[i] < b[j] {
			d = append(d, a[i])
			i++
		} else {
			d = append(d, b[j])
			j++
		}
	}
	d = append(d, a[i:len(a)]...)
	d = append(d, b[j:len(b)]...)
	return d
}

func sortIfNeeded(a []int64) []int64 {
	if sortutil.Int64sAreSorted(a) {
		return a
	}
	s := append(a[:0:0], a...)
	sortutil.Int64s(s)
	return s
}

func InBatchesInt64(input []int64, chunkSize int64) [][]int64 {
	var divided [][]int64
	inputSize := int64(len(input))
	for i := int64(0); i < inputSize; i += chunkSize {
		end := i + chunkSize
		if end > inputSize {
			end = inputSize
		}
		divided = append(divided, input[i:end])
	}
	return divided
}
