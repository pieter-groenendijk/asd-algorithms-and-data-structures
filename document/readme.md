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

```go
// Alias of a binary heap
type PriorityQueue[TPriority cmp.Ordered, TValue any] = binheap.BinHeap[TPriority, TValue]

func New[TPriority cmp.Ordered, TValue any](capacity int) *PriorityQueue[TPriority, TValue] {
	return binheap.New[TPriority, TValue](capacity)
}
```

```go
type BinHeap[TOrder cmp.Ordered, TValue any] struct {
	nodes []Node[TOrder, TValue]
}

func New[TOrder cmp.Ordered, TValue any](capacity int) *BinHeap[TOrder, TValue] {
	return &BinHeap[TOrder, TValue]{
		nodes: make([]Node[TOrder, TValue], 0, capacity),
	}
}

func (h *BinHeap[TOrder, TValue]) Push(order TOrder, value TValue) {
	h.nodes = append(h.nodes, *NewNode(order, value))

	length := len(h.nodes)
	if length == 1 {
		return
	}

	// While "heap order" incorrect, swap inserted with parent
	insertedAt := length - 1
	nodes := h.nodes
	for {
		parentAt := (insertedAt - 1) / 2 // `>> 1` is the same as `/ 2`, but more performant
		if parentAt < 0 {
			break
		}

		if order <= nodes[parentAt].order {
			break
		}

		nodes[parentAt], nodes[insertedAt] = nodes[insertedAt], nodes[parentAt]
	}
}

// for better CPU Cache performance, perhaps separate values from priorities.
// Priorities are worked on most of the time, while the value is generally
// only acted upon about once per function.
// By separating more priorities can be put into cache (spatial locality),
// while the once accessed data can remain out of cache, without any big consequence.

func (h *BinHeap[TOrder, TValue]) Pop() (TValue, bool) {
	length := len(h.nodes)
	if length == 0 {
		var value TValue
		return value, false
	}

	// Extract first while we still can
	first := h.nodes[0]

	// Move "last" to "first", and remove "last"
	h.nodes[0] = h.nodes[length-1]
	h.nodes = h.nodes[:length-1]
	length--

	// While "heap order" incorrect -> swap largest child with inserted
	nodes := h.nodes
	insertedAt := 0
	for {
		leftChildAt := insertedAt*2 + 1
		rightChildAt := insertedAt*2 + 2
		largestAt := insertedAt

		if leftChildAt < length && nodes[leftChildAt].order > nodes[largestAt].order {
			largestAt = leftChildAt
		}
		if rightChildAt < length && nodes[rightChildAt].order > nodes[largestAt].order {
			largestAt = rightChildAt
		}

		if largestAt == insertedAt { // We didn't find a larger child
			break
		}

		nodes[largestAt], nodes[insertedAt] = nodes[insertedAt], nodes[largestAt]
		insertedAt = largestAt
	}

	return first.value, true
}
```

```go
type Node[TPriority cmp.Ordered, TValue any] struct {
	order TPriority
	value TValue
}

func NewNode[TPriority cmp.Ordered, TValue any](priority TPriority, value TValue) *Node[TPriority, TValue] {
	return &Node[TPriority, TValue]{
		order: priority,
		value: value,
	}
}
```

### Description
A priority queue is implemented using a binary heap. The binary heap, although representing a binary tree, is stored
as an array, or dynamic array in this case^[A slice is the go variant]. It allows optimized priority queue
operations due to the maintenance of the shape and heap properties:

Shape: Every level is fully filled except the deepest one. By ensuring this the algorithm may have better 
existential knowledge, thus reducing checks.

Heap: In our case, every child node is lesser or equal to it's parent. The algorithm knows where the 
_largest_ node remains at, namely the root node. The heap property, combined with that the root node 
is a parent but not a child, ensures this will always remain constant. The rest of the ordering mainly
makes it easier to maintain the property itself, i.e. less comparisons for insertions and deletions.

### Performance

Operation | Best | Worst |
| ------ | --- | --- |
| Push | `O(1)` | `O(log n)` |
| Pop | `O(1)` | `O(log n)` |

#### Push
Best case occurs when the node added to the _bottom_ contains a smaller priority than it's direct parent. 
Thus not requiring no maintenance of the heap property.

Worst case occurs when the node added to the _bottom_ contains the largest priority in the collection. Thus
requiring maintenance of the heap property through all layers.

#### Pop
Best case occurs when the root node is swapped with the _last_ node containing the largest value in the
new collection. Thus requiring no maintenance of the heap property.

Worst case occurs when the root node is swapped with the _last_ node containing the smallest value in the 
new collection. Thus requiring maintenance of the heap property through all layers.

### Potential optimizations
The node structures should be implicit, not explicit; priorities should be separated from values.
Associated values are only important for the final _setting_ or _getting_, all other logic operates
on priorities. The algorithm can fit more _hot_ data in the CPU cache by minimizing on putting in
_cold_ data. Thus improving operations for a larger `n`.

A node should have more than two children. The spatial locality can be made more relevant by increasing the amount of children a parent has, since 
it leads to a more continous memory block. Downwards maintenance of the heap property ensures a comparison
per child. An iteration's comparisons may have more cache hits due to the more relevant spatial locality.

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

```go
func Search[TValue cmp.Ordered](values []TValue, value TValue) (int, error) {
	rightAt := len(values) - 1
	if rightAt < 0 {
		return -1, collections.ErrNotFound
	}

	leftAt := 0
	var middleAt int
	for leftAt != rightAt {
		middleAt = (rightAt - leftAt) / 2 + leftAt
		if value > values[middleAt] {
			leftAt = middleAt + 1
		} else { // value <= values[middleAt]
			// we assume value < values[middleAt], to prevent an repeated check, and
			// instead rely on our loop boundary condition.
			rightAt = middleAt
		}
	}

	if value == values[leftAt] {
		return leftAt, nil
	}

	return -1, collections.ErrNotFound
}
```

### Description
Binary search utilizes the ordered nature of given `values` to 
make educated guesses of where the given `value` might be. Binary research makes 
educated guesses every iteration, and checks it's relativity to the given `value`,
to determine the next educated guess. Every iteration roughly halves the search area. 

In this implementation, opposed to more traditional approaches, this _halving_ 
goes on until there is only one value left. 

This approach is taken since the actual repeating binary search, i.e. the loop, actually
uses a binary check, a `if`-`else`. Many implementations check the value at the guess `middleAt`
to be larger, smaller and equal. This way an early return may be used for equality. Yet, the 
cost is an additional conditional. The chance for a guess to be accurate is with unbounded 
length of `values` generally extremely low. Thus, this implementation just assumes for every
iteration that if `value <= values[middleAt]`, then we assume `value < values[middleAt]`, since
that is very likely to be true. Yet, since equality is theoretically possible we can't assign our
next `rightAt` to be `middleAt - 1` to prevent removing that possibility from our search area.

Then our search naturally converges, where `leftAt == rightAt`, and one value is left,
which is checked for equality to produce the desired function return types.

### Complexity
The time complexity stays the same for every non-zero input since it always keeps halving 
until there is one value left: `O(logN)`, where `N == len(values)`

### Execution time

### Potential improvements
The additional error return value, although consistent with other styles, it can be optimized
to use a `bool` instead. Go wraps interface types like `error`, thus adding interface allocation
overhead. Perhaps even easier, but arguably less safe, a int value of `-1` may be used to 
indicate that the value is not found.

The arithmatic calculation of `middleAt` could perhaps be replaced with bitwise operations, specifically
using a bitwise right shift to perform the halving division.

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

