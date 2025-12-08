package arraylist

// Golang sees array length as part of the type definition, requiring it to
// be declared as a constant. Yet, our arraylist requires it do declared at
// runtime. So, we're actually using a `slice`, which is already a dynamic
// array.
type ArrayList[TValue comparable] struct {
	space []TValue
	size int // not uint because it avoids conversions
}

func New[TValue comparable](initialCapacity int) *ArrayList[TValue] {
	return &ArrayList[TValue]{
		space: make([]TValue, initialCapacity),
		size: 0,
	}
}
