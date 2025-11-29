package maps

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"

type Map[TKey string, TValue any] interface {
	collections.Collection
	// Sets the mapping for the specified key in this map
	Set(key TKey, value TValue) 
	// Unsets the mapping for the specified key in this map
	Unset(key TKey) 
	// Gets the value for the specified key in this map
	Get(key TKey) (TValue, error)
}
