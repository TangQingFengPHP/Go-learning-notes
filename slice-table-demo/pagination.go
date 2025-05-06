package main

func FilterSlice[T any](items []T, filter func(T) bool) []T {
	var result []T
	for _, item := range items {
		if filter(item) {
			result = append(result, item)
		}
	}
	return result
}

func PaginateSlice[T any](items []T, page, pageSize int) []T {
	start := (page - 1) * pageSize
	end := start + pageSize
	if start >= len(items) {
		return []T{}
	}
	if end > len(items) {
		end = len(items)
	}
	return items[start:end]
}
