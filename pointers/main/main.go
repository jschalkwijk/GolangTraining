package main

import "fmt"

func main(){

	a := 43

	fmt.Println(a) // 43
	fmt.Println(&a) // gives the memory address where the var a is stored on the RAM or HDD/SSD


	var b *int = &a
	// the above code makes b a pointer to the memory address where an int is stored
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int

	fmt.Println(b) // gives the memory address where the var a is stored on the RAM or HDD/SSD
	fmt.Println(*b) // 43

	// this is useful
	// we can pass a memory address instead of a bunch of values (we can pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around addresses

	// everything is PASS BY VALUE in go, btw
	// when we pass a memory address, we are passing a value
}
