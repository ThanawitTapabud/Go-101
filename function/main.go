package main

import (
	"fmt"
	"math"
)

type MathFunc func(float64, float64) float64

func calculate(fn MathFunc) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(calculate(hypot))

	pos := adder()
	fmt.Println()
}
