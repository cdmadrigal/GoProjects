package main

import "fmt"

func fib(n int) int {
	a := 0
	b := 1
	var sum int
	for i := 0; i < (n / 2); i++ {
		a = a + b
		b = a + b

		}
	}

}

func main() {
	fibEven := fib(10)
	fmt.Println(fibEven)

}
