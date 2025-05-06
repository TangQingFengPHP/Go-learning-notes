package data

import "fmt"

type Person struct {
	Name string
	Age  int
}

func GenerateMockData(total int) []Person {
	people := make([]Person, total)
	for i := 0; i < total; i++ {
		people[i] = Person{
			Name: fmt.Sprintf("User%03d", i+1),
			Age:  20 + (i % 30),
		}
	}
	return people
}
