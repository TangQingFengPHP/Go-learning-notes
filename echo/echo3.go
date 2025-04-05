package main

import (
	"fmt"
	"os"
	"strings"
)

func echo3() {
	fmt.Println(os.Args[1:])
	fmt.Println(strings.Join(os.Args[1:], " , "))
}
