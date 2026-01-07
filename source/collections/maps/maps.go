package maps

type Map[TKey string, TValue any] interface {
	// Set sets the mapping for the specified key in this map
	Set(key TKey, value TValue)
	// Unset unsets the mapping for the specified key in this map
	Unset(key TKey)
	// Get gets the value for the specified key in this map
	Get(key TKey) (TValue, bool)
}
