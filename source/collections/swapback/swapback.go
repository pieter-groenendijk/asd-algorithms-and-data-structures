package swapback

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
