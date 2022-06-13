package main

import "fmt"

func main() {
	var x = -121
	fmt.Println(isPalindrome(x))
	x = 121
	fmt.Println(isPalindrome(x))
	x = 122
	fmt.Println(isPalindrome(x))

}

func isPalindrome(x int) bool {
	var y = x
	var z int
	for y != 0 {
		z *= 10
		z += y % 10
		y /= 10
	}

	return x == z
}
