package utils

// IndexOfInt returns first index of element in array
func IndexOfInt(s []int, v int) int {
	for idx, val := range s {
		if val == v {
			return idx
		}
	}
	return -1
}

// ContainsInt returns boolean which indicates if array contains value
func ContainsInt(s []int, v int) bool {
	return IndexOfInt(s, v) >= 0
}
