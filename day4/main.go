package main

import "fmt"

type ID int

var (
	b string
	c int
	d bool
	e float64
	f ID = 1
)

func main() {
	fmt.Printf("O tipo de E é %T", f)
}
