package main

import "fmt"

func main() {
	// using var
	var age int32 = 37
	const isCool = true
	var size float32 = 1.3

	// Shorthand
	// name := "Alberto"
	// email := "alberto@gmail.com"

	name, email := "Alberto", "alberto@gmail.com"

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", size)
}
