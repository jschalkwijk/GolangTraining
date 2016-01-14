package main


import "fmt"

func main() {
	basicLoop();
	nestedLoop();
	conditionsBreaksAndContinue();
}

func basicLoop(){
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

func nestedLoop(){
	// first loop runs 10 times, every one time the first loop runs,
	// the second loop runs also 10 times, each time the first loop runs
	for i := 0; i <= 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i," - ", j)
		}
	}
}

func conditionsBreaksAndContinue() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
