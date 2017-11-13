package main

import "fmt"

func numbers(n int) int {
	var x int
	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			x += i
		}
	}
	return x
}

func main() {
	sum := numbers(1000)
	fmt.Println(sum)

}
