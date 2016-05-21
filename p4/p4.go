package main

import "fmt"
import "strconv"
import "sort"

const Max = 999
const Min = 100

func main() {
	var palindromes = []int{}

	for i := Max; i >= Min; i-- {
		for j := i; j >= Min; j-- {
			p := i * j
			if isPalindrome(p) {
				palindromes = append(palindromes, p)
				break
			}
		}
	}
	sort.Ints(palindromes)
	fmt.Println(palindromes[len(palindromes)-1])
}

func isPalindrome(a int) bool {
	s := strconv.Itoa(a)
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}
