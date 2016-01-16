package main

import (
	"fmt"
)


func reverse(s string) string {
	// example: if we input the string "abcdefg"
	// we convert the string to a slice of runes. [97 98 99 100 101 102 103] (UTF-8 encoding)
	chars := []rune(s)
	fmt.Println(chars);

	// I think i and j, resemble the positions in the slice, most left, most right.
	// They switch the values, and then working towards the middle
	// in this example 0 - 6, 1 -5 , 2-4.
	// after the position switches the runes are converted back to a string.
	// I do not understand this for loop.
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		//fmt.Println(i,j)
		//fmt.Println(chars[i],chars[j]);
		//chars[i], chars[j] = chars[j], chars[i]
		fmt.Println(chars[i],chars[j]);
	}
	//fmt.Println(chars)
	return string(chars)
}

func main() {
	fmt.Printf("%v\n", reverse("abcdefg"))
}
