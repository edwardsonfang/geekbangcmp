package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 0 {
		println("Do not accept any argument")
		for i := 0; i < len(os.Args); i++ {
			fmt.Println(os.Args[i])
		}
		os.Exit(1)
	}
	println("Hello World")
}
