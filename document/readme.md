---
title: "Algorithms & Data-structures"
subtitle: "Hands-on implementation and analysis"
author: "Pieter Groenendijk (HAN 1654942)"
date: "Nov 18 2025"
lang: "en-US"
bibliography: "./references.bib"
geometry: "left=3cm,right=3cm,top=2cm,bottom=2cm"
header-includes:
    - \usepackage{attachfile}
---

\newpage
# Lists
## Dynamic array 
<!--
    - Je implementeert twee typen reeksen en maakt een vergelijking tussen de implementaties.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->

- In golang, the length of an array is part of it's type, meaning the length must be a constant. Only slices can
be created with a length specified at runtime, which already is a dynamic array...

- Not being able to specifically see if there is actually unallocated heap memory after the already defined underlying
array is something only the built-in slice implementation may know. For our implementation we don't have this granular
control and must reallocate a whole different array and copy the values upong growing.

\newpage
## Linked list
<!--
    - Je implementeert twee typen reeksen en maakt een vergelijking tussen de implementaties.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->

\newpage
## Comparison
<!--
    - Je implementeert twee typen reeksen en maakt een vergelijking tussen de implementaties.

    - Je onderbouwt de verschillen tussen de twee implementaties en maakt hierbij gebruik 
    van de concepten time complexity en execution time. Je behandelt hierbij zowel het 
    slechtst mogelijke geval als het best mogelijke geval en legt uit waarin deze van elkaar verschillen.
-->

\newpage
# Priority Queue
<!--
    - Je implementeert een priority queue.
    - Je legt uit op welke manier je de priority implementeert en onderbouwt waarom je voor deze methode hebt gekozen.
    - Je onderbouwt de performance van je implementatie en maakt hierbij gebruik van de concepten 
    time complexity en execution time. Je behandelt hierbij zowel het slechtst mogelijke geval als het 
    best mogelijke geval en legt uit waarin deze van elkaar verschillen.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->

\newpage
# Sorting
## Merge sort
<!--
    - Je implementeert een merge sort algoritme en 1 ander sorteeralgoritme naar 
    keuze (NIET parallel merge sort) en maakt een vergelijking tussen de implementaties.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->
```go
func Sort[T cmp.Ordered](list []T) []T {
	listLength := len(list)

	readFrom := list
	writeTo := make([]T, listLength)

	if listLength % 2 != 0 { // uneven
		writeTo[listLength - 1] = readFrom[listLength - 1]
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
```
### Description
Worst case time complexity: `O(n log n)`
Worst case space complexity: `O(n)`

It starts with the smallest unit of partition that is meaningful and sorted by definition: a one item partition.

Then, every iteration the whole array's partitions are _merged_ into a 
larger partitions, this is done until the whole array is sorted.
A divide-and-conquer algorithm. 

### Noteworthy optimizations
Two arrays are maintained: one for reading the state before the current iteration, and one to write new state to while in
this iteration. These are then swapped after an iteration has completed. Here, we initialize the `readFrom` as the given list,
and the `writeTo` as an new list. Lessening time spent on memory management and copying operations.

The iterative approach, opposed to resursive, also lessens time and space usage, since there is no overhead done for this
algorithm on the call stack.

### Potential improvements
Each merge iteration uses each item exactly once, and if the length of the list is larger than the CPU cache capacity, then
these are removed from it before they are used again in the next iteration. This generally leads to dreadful locality. Temporal 
locality may be improved by only merging subarrays that fit into the cache, and then merging the whole array from that point on.
Then a memory-optimized heap can be used to increase spatial locality for the last phase of mentioned improvement. [@CachePerfSorting]

Merging has significant overhead, which makes it worse for smaller partitions. By falling back to 
another sorting algorithm that can better handle this workload. Perhaps lowering time and space utilization.

Merging subarrays in reversed order would allow simpler boundary checking, perhaps lowering time and space
utilization.

## Insertion sort
<!--
    - Je implementeert een merge sort algoritme en 1 ander sorteeralgoritme naar 
    keuze (NIET parallel merge sort) en maakt een vergelijking tussen de implementaties.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->

## Comparison
<!--
    - Je implementeert een merge sort algoritme en 1 ander sorteeralgoritme naar 
    keuze (NIET parallel merge sort) en maakt een vergelijking tussen de implementaties.
    - Je onderbouwt de verschillen tussen de twee implementaties en maakt 
    hierbij gebruik van de concepten time complexity, execution time en space complexity. Je 
    behandelt hierbij zowel het slechtst mogelijke geval als het best mogelijke geval en 
    legt uit waarin deze van elkaar verschillen.
--> 

\newpage
# Hash Table
<!--
    - Je implementeert een hash table.
    - Je legt uit hoe je hashing-algoritme werkt en onderbouwt je keuze voor het specifieke algoritme.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->

\newpage
# Binary Search
<!--
    - Je implementeert een binary search.
    - Je onderbouwt de performance van je implementatie en maakt hierbij gebruik van de concepten time 
    complexity en execution time. Je behandelt hierbij zowel het slechtst mogelijke geval als het 
    best mogelijke geval en legt uit waarin deze van elkaar verschillen.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->

\newpage
# Graphs
## [Graph implementation]
<!--
    - Je implementeert Dijkstra's algoritme. Hiervoor gebruik je een door jou 
    zelf ontwikkelde implementatie van een graaf.
    - Je onderbouwt de performance van je implementatie van een graaf en maakt hierbij 
    gebruik van de concepten time complexity en execution time.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->
## Dijkstra's
<!--
    - Je implementeert Dijkstra's algoritme. Hiervoor gebruik je een door jou zelf 
    ontwikkelde implementatie van een graaf.
    - Je onderbouwt de performance van je implementatie van Dijkstra's 
    algoritme en maakt hierbij gebruik van de concepten time complexity en execution time.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->

