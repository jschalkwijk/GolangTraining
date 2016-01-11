package main

import "fmt"

	func main(){
	//	Shorthand notation

		a := 10
		b := "golang"
		c := 4.17
		d := true

		//zero value strings, the are declared but not assigned to a value
		var e int
		var f string
		var g float64
		var h bool

		// godoc.org/fmt. %v or other letters are format types.
		fmt.Println("Print Vars with format:\n")
		fmt.Printf("%v \n", a)
		fmt.Printf("%v \n", b)
		fmt.Printf("%v \n", c)
		fmt.Printf("%v \n", d)

		fmt.Println("Print Vars: \n")
		fmt.Println(a);
		fmt.Println(b);
		fmt.Println(c);
		fmt.Println(d);

		fmt.Println("Types: \n")
		fmt.Printf("%T \n", a)
		fmt.Printf("%T \n", b)
		fmt.Printf("%T \n", c)
		fmt.Printf("%T \n", d)

		fmt.Println("Print Vars zero value: \n")
		fmt.Println(e);
		fmt.Println(f);
		fmt.Println(g);
		fmt.Println(h);
	}