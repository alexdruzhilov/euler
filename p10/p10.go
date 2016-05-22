package main

import "fmt"

const N = 2000000

func main() {
	isPrime := make([]bool, N)
	for i := 0; i < N; i++ {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false

	for i := 2; i*i < N; i++ {
		if isPrime[i] {
			for j := i * i; j < N; j += i {
				isPrime[j] = false
			}
		}
	}

	sum := 0
	for i := 2; i < N; i++ {
		if isPrime[i] {
			sum += i
		}
	}

	fmt.Println(sum)
}
