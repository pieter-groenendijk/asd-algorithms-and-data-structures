package divhash

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/hashing"

func Hash(value int, size int) (int, error) {
	if size < 1 {
		return value, hashing.ErrNoHashRoom
	} 
	if value < 0 {
		return value, hashing.ErrValueIsNegative
	}

	return value % size, nil
}
