package main

import (
	"fmt"
	"os"
)

func echo() {
	fmt.Println(os.Args)
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
