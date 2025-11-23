package hashing

// Hash reduces universe of all integers down to a reasonable size.
//
// value >= 0
// size >= 1, can't return any value if ther are less than 1 options to choose from.
type Hash func(value int, size int) int


