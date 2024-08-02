package common

import "golang.org/x/exp/constraints"

// Contains checks if a slice contains a value of type T that is ordered.
func Contains[T constraints.Ordered](sourceSlice []T, value T) bool {
	for _, element := range sourceSlice {
		if element == value {
			return true
		}
	}

	return false
}
