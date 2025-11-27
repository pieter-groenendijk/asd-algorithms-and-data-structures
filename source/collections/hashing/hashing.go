package hashing

// Hash reduces universe of all integers down to the given size.
//
// value >= 0  
// 
// size >= 1, can't return any value if ther are less than 1 options to choose from.  
type Hash func(value uint32, size uint32) (uint32, error)


