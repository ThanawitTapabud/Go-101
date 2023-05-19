package main

import (
	"fmt"
	"math"
	"struct/person"
)

type Rectangle struct {
	w float64
	l float64
}

type Circle struct {
	radius float64
}

func (r Circle) area() float64 {
	return math.Pi * (math.Pow(r.radius, 2))
	
}

func main() {
	p := person.Person{
		Name: "PH4K",
		Age:  20,
	}
	fmt.Println(p)
	p.Name = "CASINO"
	fmt.Println(p)

	c := Circle{
		radius: 10,
	}
	fmt.Println(c)
	fmt.Printf("%.2f\n", c.area())
}
