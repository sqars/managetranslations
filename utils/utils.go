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

// FilterStr filters string out of an array
func FilterStr(s []string, v string) []string {
	result := []string{}
	for _, val := range s {
		if val != v {
			result = append(result, val)
		}
	}
	return result
}

// GroupByIdxValue groups nested slices by value in slice index
func GroupByIdxValue(d [][]string, idx int) map[string][][]string {
	result := map[string][][]string{}
	for _, v := range d {
		_, ok := result[v[idx]]
		if !ok {
			result[v[idx]] = [][]string{}
		}
		result[v[idx]] = append(result[v[idx]], v)
	}
	return result
}
