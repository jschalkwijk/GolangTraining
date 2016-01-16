package main
import "fmt"

func main ()  {
	foo()
	foo(1,2,3)
	foo(1,2)
	aSlice := []int{1,2,3,4}
	foo(aSlice...)
}

func foo(num ...int) []int{
	return fmt.Println(num)
}
