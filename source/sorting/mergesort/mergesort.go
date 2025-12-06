package mergesort

import (
	"cmp"
)

// Implements sorting.Sort
func Sort2[T cmp.Ordered](list []T) []T {
	listLength := len(list)
	if listLength == 0 {
		return make([]T, 0)
	}

	sortedLists := make([][]T, listLength)
	for at := 0; at < listLength; at++ {
		sortedLists[at] = list[at : at+1]
	}

	// Merge while pairs exist
	sortedLength := listLength
	for sortedLength > 1 {
		insertResultAt := 0

		// Merge each pair of sorted lists
		rightAt := 1
		for ; rightAt < sortedLength; rightAt += 2 {
			left := sortedLists[rightAt-1]
			leftLength := len(left)
			nextLeftAt := 0

			right := sortedLists[rightAt]
			rightLength := len(right)
			nextRightAt := 0

			resultLength := leftLength + rightLength
			result := make([]T, resultLength)
			nextResultAt := 0

			// insert the smallest item in the result of the two while comparison is needed
			for nextLeftAt < leftLength && nextRightAt < rightLength {
				leftItem := left[nextLeftAt]
				rightItem := right[nextRightAt]
				if leftItem < rightItem {
					result[nextResultAt] = leftItem

					nextLeftAt++
				} else {
					result[nextResultAt] = rightItem

					nextRightAt++
				}

				nextResultAt++
			}
			// insert leftover `left` items if there are any
			for ; nextLeftAt < leftLength; nextLeftAt++ {
				result[nextResultAt] = left[nextLeftAt]
				nextResultAt++
			}
			// insert leftover `right` items if there are any
			for ; nextRightAt < rightLength; nextRightAt++ {
				result[nextResultAt] = right[nextRightAt]
				nextResultAt++
			}

			sortedLists[insertResultAt] = result

			insertResultAt++
		}

		// Append leftover sorted list
		if rightAt-1 < sortedLength { // not part of a pair, happens for uneven amount of items.
			sortedLists[insertResultAt] = sortedLists[rightAt-1]
			insertResultAt++
		}

		sortedLength = insertResultAt
		sortedLists = sortedLists[0:sortedLength] // make old data invisible to algo
	}

	return sortedLists[0]
}

func Sort[T cmp.Ordered](list []T) {
	listLength := len(list)

	work := make([]T, listLength)

	partLength := 1
	mergeLength := 2
	for partLength < listLength {
		leftStartAt := 0
		rightStartAt := leftStartAt + partLength
		rightEndedAt := min(mergeLength, listLength)
		for rightStartAt < listLength {
			nextLeftAt := leftStartAt
			nextRightAt := rightStartAt
			insertAt := nextLeftAt

			// insert the smallest item in the result of the two while comparison is needed
			for nextLeftAt < rightStartAt && nextRightAt < rightEndedAt {
				leftItem := list[nextLeftAt]
				rightItem := list[nextRightAt]
				if leftItem < rightItem {
					work[insertAt] = leftItem

					nextLeftAt++
				} else {
					work[insertAt] = rightItem

					nextRightAt++
				}

				insertAt++
			}
			// insert leftover `left` items if there are any
			for ; nextLeftAt < rightStartAt; nextLeftAt++ {
				work[insertAt] = list[nextLeftAt]
				insertAt++
			}
			// insert leftover `right` items if there are any
			for ; nextRightAt < rightEndedAt; nextRightAt++ {
				work[insertAt] = list[nextRightAt]
				insertAt++
			}

			// copy from work into list
			for at := leftStartAt; at < rightEndedAt; at++ {
				list[at] = work[at]
			}

			// prepare for next iteration
			leftStartAt += mergeLength
			rightStartAt += mergeLength
			rightEndedAt = min(rightEndedAt+mergeLength, listLength)
		}

		// prepare for next merge iteration
		partLength = mergeLength
		mergeLength *= 2
	}
}
