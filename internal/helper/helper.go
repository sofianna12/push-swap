package helper

func IsSorted(vals []int) bool {
	for i := 1; i < len(vals); i++ {
		if vals[i-1] > vals[i] {
			return false
		}
	}
	return true
}
