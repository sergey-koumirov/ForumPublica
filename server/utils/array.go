package utils

// Find returns the smallest index i at which x == a[i], or len(a) if there is no such index.
func Find(a []int64, x int64) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}
