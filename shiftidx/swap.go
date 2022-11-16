// Package shiftidx is a library to modify the order of types with
// numbered indexes.
package shiftidx

import (
	"errors"
)

// Swap the values of two list positions.
func Swap[T any](list []T, i, j int) ([]T, error) {
	if i == j { // not a real swap, do nothing
		return list, nil
	}
	max := len(list)
	if i < 0 || j < 0 || max < i || max < j {
		return list, errors.New("list position(s) out of bounds")
	}

	list[i], list[j] = list[j], list[i]
	return list, nil
}
