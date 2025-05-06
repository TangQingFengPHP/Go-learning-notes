package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func GetMockUsers() []User {
	users := make([]User, 100)
	for i := range users {
		users[i] = User{
			ID:    i + 1,
			Name:  "User_" + string('A'+(i%26)),
			Email: fmt.Sprintf("user%d@example.com", i+1),
			Age:   18 + (i % 30),
		}
	}
	return users
}
