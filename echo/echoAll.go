package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echoAll1() {
	startTime := time.Now()
	fmt.Println(os.Args)
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	endTime := time.Now()
	fmt.Println("echoAll1 Elapsed time:", endTime.Sub(startTime))
	fmt.Println(s)
	echoAll2()
}

func echoAll2() {
	startTime := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " | "
	}
	endTime := time.Now()
	fmt.Println("echoAll2 Elapsed time:", endTime.Sub(startTime))
	fmt.Println(s)
	echoAll3()
}

func echoAll3() {
	startTime := time.Now()
	res := strings.Join(os.Args[:], " , ")
	endTime := time.Now()
	fmt.Println("echoAll3 Elapsed time:", endTime.Sub(startTime))
	fmt.Println(res)
	echoAll4()
}

func echoAll4() {
	var builder strings.Builder
	for i, arg := range os.Args[:] {
		if i > 0 {
			builder.WriteString(" <> ")
		}
		builder.WriteString(arg)
	}
	fmt.Println(builder.String())
}
