package main

import "fmt"

func main(){
	for i := 0; i <= 100; i++ {
		if (i % 3 == 0) && (i % 5 == 0) {
			fmt.Println(i," - FizzBuzz")
		} else if i % 5 == 0 {
			fmt.Println(i, " - Buzz")
		} else if i % 3 == 0 {
			fmt.Println(i, " - Fizz")
		} else {
			fmt.Println(i)
		}
	}

	var sum int

	for i := 0; i < 1000; i++ {

		if i % 5 == 0 {
			sum += i
		} else if i % 3 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)


}
