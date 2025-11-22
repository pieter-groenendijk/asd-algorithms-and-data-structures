package maps

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"

type Map[TKey string, TValue any] interface {
	collections.Collection
	// Sets the mapping for the specifiek key in this map
	Set(key TKey, value TValue) error
	// Unsets the mapping for the specified key in this map
	Unset(key TKey) error
	// Gets the value for the specified key in this map
	Get(key TKey) (TValue, error)
}

func New[TKey string, TValue any]() {

}
