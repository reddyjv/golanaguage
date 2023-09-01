package main

import (
	"fmt"
	"strconv"
)

func palindrome(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}
func largestpal() (int, int, int) {
	large := 0
	var mul1, mul2, product int
	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			product = i * j
			if product < large {
				break
			}
			if palindrome(product) && product > large {
				large = product
				mul1 = i
				mul2 = j

			}
		}
	}
	return large, mul1, mul2
}
func main() {
	large, n, m := largestpal()
	fmt.Println("values rae %d| %d|%d ", large, n, m)
}
