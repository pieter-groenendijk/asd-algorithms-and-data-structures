package swapback

func Pop[T any](b []T) ([]T, T, bool) {
	newLen := len(b) - 1
	var zeroValue T
	if newLen == -1 {
		return b, zeroValue, false
	}

	value := b[newLen]
	b[newLen] = zeroValue

	return b[:newLen], value, true
}

func Remove[T any](b []T, index int) []T {
	newLen := len(b) - 1
	if newLen == -1 {
		return b
	}

	b[index] = b[newLen]
	var zeroValue T
	b[newLen] = zeroValue // May help relieve some memory

	// TODO: In some cases reallocate to smaller size

	return b[:newLen]
}
