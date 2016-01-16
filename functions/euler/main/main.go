package main

import (
	"fmt"
	"strconv"
)


/*
Largest palindrome product
Problem 4
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
func main()  {
	palindrome()
}

func palindrome(){
	var biggest int
	var sum string
	// create empty slice assigned to numbers
	var numbers = make([]int,0)

	// append integers to slice stored in the var numbers.
	for i := 100; i<= 999; i++{
		numbers = append(numbers,i)
	}
	// for every value in the numbers slice we have to multiply the value
	// of the slice * all the numbers from the numbers slice.
	for _, value := range numbers{

		for i := 0; i < len(numbers); i++{
			output := value * numbers[i]
			// the reverse function reverses strings only
			// use the build in module strconv to convert output to a string
			t := strconv.Itoa(output)
			// check if the string is equal to the string in reverse
			// if true, check if the biggest so far, is smaller. if it is,
			// the current number is the biggest and we set the the current output to biggest.
			// which is an INT (so not the converted string value)
			// in Sum we store the , drum-role..,  the sum which created the largest number
			if t == reverse(t) {
				//fmt.Println(t) uncomment to see al the palindrome numbers
				if biggest < output {
					biggest = output
					sum = strconv.Itoa(value) + "*"+ strconv.Itoa(numbers[i])
				}
			}
		}

	}

	// print out the biggest Palindrome number.
	fmt.Println("This sum: ", sum, " creates the largest palindrome made from the product of two 3-digit numbers.")
	fmt.Println(biggest)

	reverse("hello")

}

// I can't take credit for this reverse function but I will try to explain the functionality.
// source: http://golangcookbook.com/chapters/strings/reverse/

/* Reverse takes a value of type string which is assigned to a variable s, it returns a string*/

func reverse(s string) string {
	// example: if we input the string "hello"
	// we convert the string to a slice of runes. [106 111 114 110] (UTF-8)
	chars := []rune(s)
	//fmt.Println(chars);
	//fmt.Println(len(chars))
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		//fmt.Println(chars[i],chars[j]);
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

