package main

import "fmt"

func main() {

	var i = 15.5
	var s = "Hi There!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%v#\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%%\n", i)

	fmt.Printf("%v\n", s)
	fmt.Printf("%v#\n", s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%%\n", s)

}