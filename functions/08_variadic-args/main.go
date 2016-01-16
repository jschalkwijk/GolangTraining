package main

import "fmt"

func main() {
	// dit is een list of SLice.

	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data...)
	// je kan die doorgeven aan de average functie door data... te doen, dan pakt hij de
	// waardes een voor een.
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
