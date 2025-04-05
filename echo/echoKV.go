package main

import (
	"fmt"
	"os"
	"strconv"
)

func echoKV1() {
	var s string
	for i := 0; i < len(os.Args); i++ {
		s += "key=" + strconv.Itoa(i) + ":KV1" + " value=" + os.Args[i] + "\n"
	}
	fmt.Println(s)
	echoKV2()
}

func echoKV2() {
	s := ""
	for i, arg := range os.Args[:] {
		s += "key=" + strconv.Itoa(i) + ":KV2" + " value=" + arg + "\n"
	}
	fmt.Println(s)
}
