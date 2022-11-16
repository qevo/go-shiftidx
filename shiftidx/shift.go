// Package shiftidx is a library to modify the order of types with
// numbered indexes.
package shiftidx

// Move element positions based on a function.
func Custom[T any](list []T, fn func([]T) ([]T, error)) ([]T, error) {
	return fn(list)
}

// Move all element positions to the other half of the slice,
// maintaining relative distance from the middle.
func Reverse[T any](list []T) []T {
	var div int
	cap := len(list)
	if cap%2 == 0 {
		div = (cap / 2)
	} else {
		div = (cap - 1) / 2
	}

	for i := 0; i < div; i++ {
		list, _ = Swap(list, i, cap-1-i)
	}

	return list
}

// Move all element positions based on an offset amount.
func Rotate[T any](list []T, n int) []T {
	n = parseOffset(n, len(list))
	// no need to actually shift when offset is 0
	if n == 0 {
		return list
	}

	cap := len(list)
	div := cap - n
	y := list[div:]
	y = append(y, list[:div]...)
	list = y

	return list
}

// Normalize the given offset.
// - Wrap any offset greater than character count.
// - Invert negative offset value.
func parseOffset(o, l int) int {
	if o < 0 && o*-1 > l {
		o = ((o * -1) % l) * -1 // wrap large negative offset
	} else if o > l {
		o = o % l // wrap large positive offset
	}

	if o < 0 {
		o = o + l // invert negative offset
	}

	return o
}
