package main

import "math"
import "fmt"

const Index = 10001

func main() {
	i := 2
	n := 1

	for n < Index {
		i += 1

		if remainder(i, 2) != 0 && isPrime(i) {
			n++
		}
	}
	fmt.Println(i)
}

func isPrime(x int) bool {
	for i := 3; i <= sqrt(x); i++ {
		if remainder(x, i) == 0 {
			return false
		}
	}
	return true
}

func sqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

func remainder(x int, y int) float64 {
	return math.Remainder(float64(x), float64(y))
}
