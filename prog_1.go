package main

import (
	"fmt"

	"github.com/AlexV656/Go/prog_2"
)

var ExportedVariable string = "Hello, World!"

func main() {
	var msg string = "Hello"
	fmt.Println(ExportedVariable)
	fmt.Println(prog_2.ExportedVariable)
	for i := 1; i < 10; i++ {
		fmt.Println(msg)
	}
}
