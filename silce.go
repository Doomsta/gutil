package gutil

import "sort"

func Page[T any](list []T, start, length int, less func(i, j int) bool) []T {
	if less != nil {
		sort.Slice(list, less)
	}

	n := len(list)
	if length == 0 {
		length = n - start
	}

	end := Clamp(0, start+length, n)
	start = Clamp(0, start, end)

	return list[start:end]
}
