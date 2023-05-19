package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum++
	}
	fmt.Println(sum)

	if sum < 10 {
		fmt.Println("less than 10")
	} else {
		fmt.Println("= or > 10")
	}
}
