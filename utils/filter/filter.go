package filter

func Filter[V any](arr []V, predicate func(item V, idx int) bool) []V {
	var result []V

	for i, item := range arr {
		if predicate(item, i) {
			result = append(result, item)
		}
	}

	return result
}
