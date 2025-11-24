package divhash

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/hashing"

func Hash(value uint32, size uint32) (uint32, error) {
	if size < 1 {
		return 0, hashing.ErrNoHashRoom
	} 

	return value % size, nil
}
