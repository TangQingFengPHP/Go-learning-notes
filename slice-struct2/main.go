package main

import (
	"fmt"
	"sort"
)

// Person 定义结构体
type Person struct {
	Name string
	Age  int
}

// String 结构体方法（值接收者）
func (p Person) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

// Grow 结构体方法（指针接收者）
func (p *Person) Grow() {
	p.Age++
}

func main() {
	// 创建结构体切片
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}

	// 遍历切片（值传递）
	for _, p := range people {
		fmt.Println(p)
	}

	// 修改结构体（需使用指针切片）
	peoplePtr := []*Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}

	for _, p := range peoplePtr {
		p.Grow() // 调用指针接收者方法
	}

	fmt.Println(peoplePtr)

	// 按年龄排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sorted people:", people)

	// 动态添加元素
	people = append(people, Person{"David", 40})
	fmt.Println("Updated people:", people)
}
