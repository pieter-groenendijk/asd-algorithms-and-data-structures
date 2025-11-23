package hashing

import "errors"

var ErrNoHashRoom = errors.New("Hash function has no room, size is smaller than 1.")
var ErrValueIsNegative = errors.New("Hash function can't hash negative values.")
