package main

import "fmt"

const p = "death & taxes"
const q string = "death & taxes"

// Iota is a constant that incrementes itself when asigned to multiple values.
// if assigned again count resets.
const (
A = iota //0
	B = iota //1
	C = iota //2
)

const (
	D = iota //0
	E = iota //1
	F = iota //2
)

const (
	_ = iota //0
	H = iota * 10 // 1 * 10
	I = iota * 10 // 2 * 10
)

func main(){

	const r = "death & taxes"
	const s string = "death & taxes"
	const t = 42

	fmt.Printf("%p - ", p)
	fmt.Println("\n")
	fmt.Printf("%q - ", t)
	fmt.Println("\n")

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
}
