package main

import "fmt"

func main() {
	sum := 0
	sumOfSquares := 0
	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
		sum += i
	}
	fmt.Println(sum*sum - sumOfSquares)
}
