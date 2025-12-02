package sorting

import "cmp"

type SortInPlaceFunc[T cmp.Ordered] func(list []T) 

type SortFunc[T cmp.Ordered] func(list []T) []T
