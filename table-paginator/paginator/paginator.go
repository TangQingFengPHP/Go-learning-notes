package paginator

type Page[T any] struct {
	Items      []T
	PageNo     int
	PageSize   int
	TotalItems int
}

func Paginate[T any](data []T, pageNo, pageSize int) Page[T] {
	total := len(data)
	start := (pageNo - 1) * pageSize
	end := start + pageSize
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	return Page[T]{
		Items:      data[start:end],
		PageNo:     pageNo,
		PageSize:   pageSize,
		TotalItems: total,
	}
}
