package main

import "fmt"

// Filter 泛型过滤函数（过滤满足条件的元素）
func Filter[T any](slice []T, test func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if test(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map 泛型映射函数（转换元素类型）
func Map[T any, U any](slice []T, mapper func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = mapper(v)
	}
	return result
}

// Repository 泛型结构体（存储切片）
type Repository[T any] struct {
	Data []T
}

// Add 向存储切片中添加元素
func (r *Repository[T]) Add(item T) {
	r.Data = append(r.Data, item)
}

func main() {
	// 过滤整数切片
	numbers := []int{1, 2, 3, 4, 5}
	evenNumbers := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Even numbers:", evenNumbers)

	// 转换字符串切片为长度切片
	words := []string{"apple", "banana", "cherry"}
	lengths := Map(words, func(s string) int {
		return len(s)
	})
	fmt.Println("Word lengths:", lengths)

	// 使用泛型结构体
	repo := Repository[string]{Data: []string{"Go", "Rust"}}
	repo.Add("C++")
	fmt.Println("Repository:", repo.Data)
}
