package main

import "fmt"

func main(){
	fmt.Println(biggestNum(10,20,9,60,5,17))
}

func biggestNum(num ...int) int{
	var biggest int
	for _, value := range num {
		if value > biggest {
			biggest = value
		}
	}
	return biggest
}