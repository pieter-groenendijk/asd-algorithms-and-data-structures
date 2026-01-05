package swapback

type SwapbackArray[T any] struct {
	values []T
}

func (a *SwapbackArray[T]) Get(index int) T {
	return a.values[index]
}

func (a *SwapbackArray[T]) Set(index int, value T) {
	a.values[index] = value
}

func (a *SwapbackArray[T]) Remove(index int) {
	newLen := len(a.values) - 1
	if newLen == 0 {
		return
	}

	a.values[index] = a.values[newLen]
	var zeroValue T
	a.values[newLen] = zeroValue // May help relieve some memory
	a.values = a.values[:newLen]
}
