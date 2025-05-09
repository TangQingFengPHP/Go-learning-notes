package main

import "fmt"

type User struct {
	ID   int
	Name string
}

// Deduplicate 切片去重
func Deduplicate[T comparable](s []T) []T {
	seen := make(map[T]bool)
	var result []T
	for _, v := range s {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// SliceToMap 切片转map
func SliceToMap[T any, K comparable](s []T, keyFunc func(T) K) map[K]T {
	m := make(map[K]T)
	for _, item := range s {
		m[keyFunc(item)] = item
	}
	return m
}

type Identifiable interface {
	GetID() int
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

func (p Product) GetID() int {
	return p.ID
}

func FindByID[T Identifiable](items []T, id int) (T, bool) {
	for _, item := range items {
		if item.GetID() == id {
			return item, true
		}
	}
	var zero T
	return zero, false
}

func main() {
	// 去重
	nums := []int{1, 2, 2, 3}
	strs := []string{"a", "a", "b"}
	fmt.Println(Deduplicate(nums))
	fmt.Println(Deduplicate(strs))

	// 切片转map
	users := []User{
		{1, "John"},
		{2, "Jane"},
	}
	userMap := SliceToMap[User, int](users, func(u User) int {
		return u.ID
	})
	fmt.Println(userMap)

	products := []Product{
		{1, "Laptop", 999.9},
	}
	result, found := FindByID(products, 1)

	// ID 找不到就返回空切片
	result2, found2 := FindByID(products, 2)
	fmt.Println(result)
	fmt.Println(found)
	fmt.Println(result2) // {0 0}，中间是空字符串
	fmt.Println(found2)

}
