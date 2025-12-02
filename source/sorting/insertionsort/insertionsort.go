package insertionsort

import "cmp"

// Implements `sorting.SortInPlace`
func Sort[T cmp.Ordered](list []T) {
	// lookAt <- 0
	// currentItem <- list[lookAt]

	// compareWith <- lookAt - 1
	// for: compareWith > 0; compareWith--
		// compareWithItem <- list[compareWith]
		// if currentItem < compareWithItem 
			// list[compareWith + 1] <- compareWithItem
		// else 
			// break

	// list[compareWith + 1] = currentItem
	listLength := len(list)
	for lookAt := 0; lookAt < listLength; lookAt++ {
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
