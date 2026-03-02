// Package helper provides utility helper functions used across
// the push-swap project.
package helper

// IsSorted reports whether the given slice of integers is sorted in
// strictly ascending order from index 0 to the last element.
//
// It returns true for an empty or single-element slice.
//
// Parameters:
//   - vals: the slice to check, where index 0 represents the top of the stack.
//
// Returns:
//   - true if the slice is sorted in ascending order
//   - false otherwise

func IsSorted(vals []int) bool {
	for i := 1; i < len(vals); i++ {
		if vals[i-1] > vals[i] {
			return false
		}
	}
	return true
}
