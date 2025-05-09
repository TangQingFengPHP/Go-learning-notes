package main

import (
	"fmt"
	"slices"
)

func main() {
	// 切片定义，此时是nil切片，长度和容量都为0
	var s []int
	fmt.Println(s, s == nil, len(s), cap(s)) // [] true 0 0

	// 空切片
	s = []int{}                              // 已初始化但无元素
	fmt.Println(s, s == nil, len(s), cap(s)) // [] false 0 0

	// 切片创建
	// 使用字面量
	s = []int{1, 2, 3}
	fmt.Println(s, len(s), cap(s)) // [1 2 3] 3 3

	// 从数组或切片切割
	arr := [5]int{0, 1, 2, 3, 4}
	s2 := arr[1:4]
	fmt.Println(arr, s2, len(s2), cap(s2)) // [0 1 2 3 4] [1 2 3] 3 4

	// 使用make创建切片
	s3 := make([]int, 3)
	fmt.Println(s3, s3 == nil, len(s3), cap(s3)) // [0 0 0] false 3 3
	s4 := make([]int, 3, 5)
	fmt.Println(s4, s4 == nil, len(s4), cap(s4)) // [0 0 0] false 3 5

	// 切片的切割和操作
	s5 := []int{10, 20, 30, 40, 50}
	fmt.Println(s5[1:4]) // [20 30 40]
	fmt.Println(s5[:3])  // [10 20 30]
	fmt.Println(s5[2:])  // [30 40 50]
	fmt.Println(s5[:])   // [10 20 30 40 50]

	// 删除元素
	s6 := []int{1, 2, 3, 4}
	s6 = append(s6[:2], s6[3:]...)
	fmt.Println(s6, len(s6), cap(s6)) // [1 2 4] 3 4

	// 清空切片
	// 长度置为0，保留容量
	s6 = s6[:0]
	fmt.Println(s6, s6 == nil, len(s6), cap(s6))
	// 释放底层数组
	s6 = nil
	fmt.Println(s6, s6 == nil, len(s6), cap(s6))

	// 遍历切片
	// for 循环遍历
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Using for loop")
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for...range 循环遍历
	fmt.Println("Using for...range loop")
	for i, v := range slice {
		fmt.Println(i, v)
	}

	// 高级技巧
	bigSlice := make([]int, 1000000)
	smallPart := make([]int, 10)
	copy(smallPart, bigSlice[:10])

	// 切片配合结构体
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	for _, user := range users {
		fmt.Println(user)
		fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Name)
	}

	// 切片配合接口
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 5},
	}
	printAreas(shapes)

	// 切片配合泛型
	ints := []int{1, 2, 3, 4, 5}
	strs := []string{"Go", "Java", "Python"}

	PrintSlice[int](ints)
	PrintSlice[string](strs)
	// 泛型参数可传可不传，可以自动推断
	PrintSlice(ints)
	PrintSlice(strs)

	// 切片 + 泛型 + 结构体
	filtered := FilterSlice(users, func(u User) bool { return u.ID%2 == 1 })
	for _, u := range filtered {
		fmt.Println(u)
	}

	// 切片 + 泛型 + 排序
	people := []Person{
		// 指定键名是为了顺序不论
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 50},
		{Name: "Charlie", Age: 30},
		// 可以不用指定键名
		{"Eve", 31},
	}
	slices.SortFunc(people, func(a, b Person) int {
		return a.Age - b.Age
	})
	fmt.Println(people)

	// slices包常用函数
	ages := []int{32, 18, 45, 21, 18}

	// 排序
	slices.Sort(ages)
	fmt.Println("Sorted:", ages)

	// 查找
	idx := slices.Index(ages, 45)
	fmt.Println("Index of 45:", idx)

	// 是否包含
	fmt.Println("Contains 21:", slices.Contains(ages, 21))

	// 去重（只去相邻重复）
	dedup := slices.Clone([]int{1, 1, 2, 2, 2, 3})
	slices.Compact(dedup)
	fmt.Println("Compact:", dedup)

	// 插入
	inserted := slices.Insert([]int{1, 2, 4}, 2, 3)
	fmt.Println("After insert:", inserted)

	// 删除
	deleted := slices.Delete(inserted, 1, 3)
	fmt.Println("After delete:", deleted)
	fmt.Println("Original:", inserted)
}

type Person struct {
	Name string
	Age  int
}

type User struct {
	ID   int
	Name string
}

type Shape interface {
	Area() float64
}

type Circle struct{ Radius float64 }
type Rectangle struct{ Width, Height float64 }

func (circle Circle) Area() float64 { return 3.14 * circle.Radius * circle.Radius }
func (r Rectangle) Area() float64   { return r.Width * r.Height }

func printAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
}

// PrintSlice 定义泛型函数，打印任何类型的切片
func PrintSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// FilterSlice 定义泛型函数，过滤切片元素
func FilterSlice[T any](items []T, filter func(T) bool) []T {
	var result []T
	for _, item := range items {
		if filter(item) {
			result = append(result, item)
		}
	}
	return result
}
