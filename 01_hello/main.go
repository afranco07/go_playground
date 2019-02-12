package main

import "fmt"
import "os"

func main() {
	args := os.Args
	minArgs := os.Args[1:]
	fmt.Println("Hello World", args, minArgs)
}
