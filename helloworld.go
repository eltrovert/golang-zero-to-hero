package main

import "fmt"

func main() {
	// This is an example of inline comment
	// Print "Hello World"
	fmt.Println("Hello World")
	// fmt.Println("This would not be printed")

	/*
		This is an example of  multiline comment
		It can be used to create a documentation
	*/
	fmt.Println("Hello", "World", "!")
	fmt.Println("Hello", "World", 2)

	/*
		fmt.Println("Not", "Printed", "!")
		fmt.Println("Also not printed")
	*/
}
