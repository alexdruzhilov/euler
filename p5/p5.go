package main

import "fmt"
import "math"

func main() {
	var x = []int{1}

	for i := 2; i <= 20; i++ {
		k := i
		for j := 0; j < len(x); j++ {
			if math.Remainder(float64(k), float64(x[j])) == 0 {
				k = k / x[j]
			}
		}
		x = append(x, k)
	}

	s := 1
	for i := 0; i < len(x); i++ {
		s = s * x[i]
	}

	fmt.Println(s)
}
