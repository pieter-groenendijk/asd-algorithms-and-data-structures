package mergesort

import (
	"cmp"
)

func Sort[T cmp.Ordered](list []T) []T {
	listLength := len(list)

	readFrom := list
	writeTo := make([]T, listLength)

	if listLength%2 != 0 { // uneven
		writeTo[listLength-1] = readFrom[listLength-1]
	}

	var leftStartAt int
	var rightStartAt int
	var rightEndedAt int

	var nextLeftAt int
	var nextRightAt int
	var insertAt int

	partLength := 1
	mergeLength := 2
	for partLength < listLength {
		leftStartAt = 0
		rightStartAt = leftStartAt + partLength
		rightEndedAt = min(mergeLength, listLength)
		insertAt = 0
		for rightStartAt < listLength {
			nextLeftAt = leftStartAt
			nextRightAt = rightStartAt

			// merge items
			for nextLeftAt < rightStartAt && nextRightAt < rightEndedAt {
				if readFrom[nextLeftAt] < readFrom[nextRightAt] {
					writeTo[insertAt] = readFrom[nextLeftAt]

					nextLeftAt++
				} else {
					writeTo[insertAt] = readFrom[nextRightAt]

					nextRightAt++
				}

				insertAt++
			}
			// insert leftover `left` items if there are any
			for ; nextLeftAt < rightStartAt; nextLeftAt++ {
				writeTo[insertAt] = readFrom[nextLeftAt]
				insertAt++
			}
			// insert leftover `right` items if there are any
			for ; nextRightAt < rightEndedAt; nextRightAt++ {
				writeTo[insertAt] = readFrom[nextRightAt]
				insertAt++
			}

			// prepare for next iteration
			leftStartAt += mergeLength
			rightStartAt += mergeLength
			rightEndedAt = min(rightEndedAt+mergeLength, listLength)
		}

		// prepare for next merge iteration
		partLength += partLength
		mergeLength += mergeLength

		readFrom, writeTo = writeTo, readFrom
	}

	return readFrom
}
