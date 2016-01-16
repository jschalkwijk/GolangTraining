package main
import "fmt"


func main(){
	fmt.Println(multiReturn(5))
}

/*
func multiReturn(i int) (int,bool){
	var t bool
	if i%2 == 0 {
		 t = true
	 } else {
		 t = false
	 }
	 return i/2, t
}*/

func multiReturn(i int) (float64,bool){
	return float64(i) / 2, i%2 == 0
}
