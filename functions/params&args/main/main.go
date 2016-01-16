package main
import "fmt"

func main(){
	fmt.Println(greet("Jorn ","Schalkwijk"))
	fmt.Println(greet2("Jorn ","Schalkwijk"))
	fmt.Println(greet3("Jorn ","Schalkwijk"))
}

func greet(fname,lname string) string {
	return fmt.Sprint(fname,lname)
}

func greet2(fname,lname string) (s string) {
	s = fmt.Sprint(fname,lname)
	return
}

func greet3(fname,lname string) (string, string) {
	return fmt.Sprint(fname,lname), fmt.Sprint(lname,fname)
}