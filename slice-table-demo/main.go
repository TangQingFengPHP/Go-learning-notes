package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type PageData struct {
	Users    []User
	Page     int
	PrevPage int
	NextPage int
}

func main() {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pageSize := 10
		pageStr := r.URL.Query().Get("page")
		page, _ := strconv.Atoi(pageStr)
		if page <= 0 {
			page = 1
		}

		users := GetMockUsers()

		filtered := FilterSlice(users, func(u User) bool {
			return u.Age > 25
		})

		pagedUsers := PaginateSlice(filtered, page, pageSize)

		data := PageData{
			Users:    pagedUsers,
			Page:     page,
			PrevPage: max(1, page-1),
			NextPage: page + 1,
		}

		tmpl.Execute(w, data)
	})

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
