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
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3
	fmt.Print(len(meuArray))

	for i, v := range meuArray {
		fmt.Printf("\nO valor de é %d e o valor é %d\n", i, v)
	}
}
