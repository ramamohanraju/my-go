package main

import "fmt"

func main() {
	// print hello world
	println("Hello, World!")
	test := 1
	println(test)
	fmt.Printf("test address: %d", &test)
	//fmt.Printf("test address: %d", *&test)

}
