package insertionsort

import "cmp"

// Implements sorting.SortInPlace
func Sort[T cmp.Ordered](list []T) {
	listLength := len(list)
	for lookAt := 1; lookAt < listLength; lookAt++ {
		currentItem := list[lookAt]

		compareWith := lookAt - 1
		for ; compareWith >= 0; compareWith-- {
			compareWithItem := list[compareWith]
			if currentItem < compareWithItem {
				list[compareWith + 1] = compareWithItem
			} else {
				break
			}
		}

		list[compareWith + 1] = currentItem
	}
}
