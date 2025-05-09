package main

import "fmt"

// IDisplay 定义接口
type IDisplay interface {
	Display() string
}

// Product 实现接口的结构体
type Product struct {
	Name  string
	Price float64
}

func (p Product) Display() string {
	return fmt.Sprintf("%s ($%.2f)", p.Name, p.Price)
}

type User struct {
	Username string
	Email    string
}

func (u User) Display() string {
	return fmt.Sprintf("%s <%s>", u.Username, u.Email)
}

type DisplayBox[T IDisplay] struct {
	Items []T
}

func (b *DisplayBox[T]) Add(item T) {
	b.Items = append(b.Items, item)
}

func (b *DisplayBox[T]) ShowAll() {
	for _, item := range b.Items {
		fmt.Println(item.Display())
	}
}

func main() {
	box := DisplayBox[IDisplay]{}
	box.Add(Product{Name: "Laptop", Price: 999.99})
	box.Add(User{"Alice", "alice@example.com"})

	box.ShowAll()
}
