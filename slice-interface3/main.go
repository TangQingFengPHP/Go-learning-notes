package main

import "fmt"

type Dog struct {
	Breed string
}

// ToInterfaceSlice converts a slice of type T to a slice of type interface{}
func ToInterfaceSlice[T any](s []T) []any {
	result := make([]any, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}

func main() {
	dogs := []Dog{
		{Breed: "Husky"},
		{Breed: "Poodle"},
	}
	anySlice := ToInterfaceSlice(dogs)
	fmt.Println(anySlice)
}
