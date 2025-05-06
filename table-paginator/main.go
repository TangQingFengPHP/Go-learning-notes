package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"table-paginator/data"
	"table-paginator/paginator"
)

func main() {
	people := data.GenerateMockData(100)
	pageSize := 10
	pageNo := 1
	reader := bufio.NewReader(os.Stdin)

	for {
		page := paginator.Paginate(people, pageNo, pageSize)
		fmt.Printf("\n[Page %d/%d] Showing %d of %d rows:\n",
			page.PageNo,
			(page.TotalItems+pageSize-1)/pageSize,
			len(page.Items),
			page.TotalItems)

		for i, p := range page.Items {
			fmt.Printf("%2d. Name: %-10s Age: %d\n", (page.PageNo-1)*pageSize+i+1, p.Name, p.Age)
		}

		fmt.Print("\n[n] Next page | [p] Previous page | [g 5] Go to page 5 | [q] Quit\n> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch {
		case input == "n":
			pageNo++
		case input == "p":
			if pageNo > 1 {
				pageNo--
			}
		case strings.HasPrefix(input, "g "):
			var newPage int
			fmt.Sscanf(input, "g %d", &newPage)
			if newPage >= 1 {
				pageNo = newPage
			}
		case input == "q":
			return
		default:
			fmt.Println("Invalid input. Please try again.")
		}
	}

}
