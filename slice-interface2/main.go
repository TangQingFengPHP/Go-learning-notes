package main

import (
	"fmt"
	"math"
)

// Shape 定义接口
type Shape interface {
	Area() float64
}

// Circle 定义圆形结构体（实现该接口）
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Rectangle 定义矩形结构体（实现该接口）
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// 创建接口类型切片
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 5},
	}

	// 统一计算面积
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
	}
	fmt.Printf("总面积是：%.2f\n", totalArea)

	// 类型断言获取具体类型
	for _, s := range shapes {
		switch v := s.(type) {
		case Circle:
			fmt.Printf("Circle with radius %.1f\n", v.Radius)
		case Rectangle:
			fmt.Printf("Rectangle %.1fx%.1f\n", v.Width, v.Height)
		}
	}
}
