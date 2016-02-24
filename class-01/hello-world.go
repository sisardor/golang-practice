package main

import (
	"fmt"
	"time"
	//"io/ioutil"
	//m "math"
	//"net/http"
	//"strconv"
)

// fib returns a function that returns 
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func beyondHello() {
	var x int
	x = 3
	y := 4
	sum, prod := learnMultiple(x ,y)
	fmt.Println("sum: ", sum," \tprod: ", prod)
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func main() {
	fmt.Println("Hello World!")

	fmt.Println("Running Fibonacci")
	f := fib()

	// Function calls are evaluated left-to-right
	fmt.Println( f(), f(), f(), f() )
	fmt.Print("Current time ")
	s := fmt.Sprintf("`%s`", time.Now())
	fmt.Println(s)

	fmt.Println("\nbeyondHello()")
	beyondHello()
}


