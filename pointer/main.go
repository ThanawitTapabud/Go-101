package main

import "fmt"

func Double(n *int) {
	*n *= 2
}

func AppendHello(x *string) {
	*x = "Hello, " + *x
}
func main() {
	var s string
	var p *string

	p = &s //& = tell address of s
	*p = "Hello"
	fmt.Println(s, p)

	m := 4
	Double(&m)
	fmt.Println(m)
	name := "PH4K"
	AppendHello(&name)
	fmt.Println(name)
}
