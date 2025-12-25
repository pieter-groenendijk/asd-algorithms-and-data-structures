package mergesort

import (
	"cmp"
)

func Sort[T cmp.Ordered](list []T) []T {
	listLength := len(list)
	if listLength <= 1 {
		return list
	}

	readFrom := list
	writeTo := make([]T, listLength)
	writeTo[listLength - 1] = readFrom[listLength - 1]

	pairSize := 2
	pyramidSize := 4
	for pairSize <= listLength { // horizontal iterations
		// We mergesort every EVEN pair LEFT-TO-RIGHT
		leftAt := 0
		rightAt := pairSize - 1
		writeAt := leftAt
		for rightAt < listLength {
			// start merging this pair
			for leftAt != rightAt {
				if readFrom[leftAt] < readFrom[rightAt] {
					writeTo[writeAt] = readFrom[leftAt]
					leftAt++
				} else {
					writeTo[writeAt] = readFrom[rightAt]
					rightAt--
				}

				writeAt++
			}
			writeTo[writeAt] = readFrom[leftAt]

			// already move to next iteration's pair
			leftAt += pyramidSize
			rightAt += pyramidSize
		}


		// We mergesort every UNEVEN pair RIGHT-TO-LEFT
		rightAt = pyramidSize - 1
		leftAt = rightAt - pyramidSize + 1
		writeAt = rightAt
		for rightAt < listLength {
			// start merging this pair
			for leftAt != rightAt {
				if readFrom[leftAt] < readFrom[rightAt] {
					writeTo[writeAt] = readFrom[leftAt]
					leftAt++
				} else {
					writeTo[writeAt] = readFrom[rightAt]
					rightAt--
				}

				writeAt++
			}
			writeTo[writeAt] = readFrom[leftAt]

			// already move to next iteration's pair
			leftAt += pyramidSize
			rightAt += pyramidSize
		}

		// already prepare for next iteration
		pairSize += pairSize
		pyramidSize += pyramidSize
		writeTo, readFrom = readFrom, writeTo
	}

	// If it's the case, add the remaining item which could not be paired in the first horizontal iteration

	return readFrom

	// Horizontal linear iteration of 
		// We sort every even pair from left to right
			// merge as said
		// We sort every uneven pair from right to left
			// merge as said



	// var leftStartAt int
	// var rightStartAt int
	// var rightEndedAt int 

	// var nextLeftAt int
	// var nextRightAt int
	// var insertAt int

	// partLength := 1
	// mergeLength := 2
	// for partLength < listLength {
	// 	leftStartAt = 0
	// 	rightStartAt = leftStartAt + partLength
	// 	rightEndedAt = min(mergeLength, listLength)
	// 	insertAt = 0
	// 	for rightStartAt < listLength {
	// 		nextLeftAt = leftStartAt
	// 		nextRightAt = rightStartAt

	// 		// merge items
	// 		for nextLeftAt < rightStartAt && nextRightAt < rightEndedAt {
	// 			if readFrom[nextLeftAt] < readFrom[nextRightAt] {
	// 				writeTo[insertAt] = readFrom[nextLeftAt]

	// 				nextLeftAt++
	// 			} else {
	// 				writeTo[insertAt] = readFrom[nextRightAt]

	// 				nextRightAt++
	// 			}

	// 			insertAt++
	// 		}
	// 		// insert leftover `left` items if there are any
	// 		for ; nextLeftAt < rightStartAt; nextLeftAt++ {
	// 			writeTo[insertAt] = readFrom[nextLeftAt]
	// 			insertAt++
	// 		}
	// 		// insert leftover `right` items if there are any
	// 		for ; nextRightAt < rightEndedAt; nextRightAt++ {
	// 			writeTo[insertAt] = readFrom[nextRightAt]
	// 			insertAt++
	// 		}

	// 		// prepare for next iteration
	// 		leftStartAt += mergeLength
	// 		rightStartAt += mergeLength
	// 		rightEndedAt = min(rightEndedAt+mergeLength, listLength)
	// 	}

	// 	// prepare for next merge iteration
	// 	partLength += partLength
	// 	mergeLength += mergeLength

	// 	readFrom, writeTo = writeTo, readFrom
	// }

	// return readFrom
}


