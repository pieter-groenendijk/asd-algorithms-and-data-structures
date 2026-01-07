package arraylist

type ArrayList[TValue comparable] struct {
	space  []TValue
	length int // not uint because it avoids conversions
}

func New[TValue comparable](initialCapacity int) *ArrayList[TValue] {
	return &ArrayList[TValue]{
		space:  make([]TValue, initialCapacity),
		length: 0,
	}
}
