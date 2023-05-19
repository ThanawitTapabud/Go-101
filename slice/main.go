package main

import "fmt"

func main() {
	// Array - im
	var a [10]int
	a[0] = 9
	a[1] = 8

	fmt.Println(a)

	var b [2]string = [2]string{"YEE", "HELL"}

	fmt.Println(b)

	// Slice mutable (unfixed)
	var c []int
	c = append(c, 2)
	c = append(c, 3)
	fmt.Println(c)

	var d = []string{
		"Nokia",
		"Apple",
		"Red Magic",
	}
	d = append(d, "Vivo")
	d = append(d, "ROG Phone")
	fmt.Println(d[0:2]) // index 0 to 2
	fmt.Println(d[1:4]) // index 1 to 4

}
