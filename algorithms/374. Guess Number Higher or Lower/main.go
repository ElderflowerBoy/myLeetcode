package main

import "fmt"

const target = 5

func main() {
	fmt.Println(guessNumber(11))
}

func guessNumber(n int) int {
	min := 1
	max := n

	for min <= max {
		avg := (min + max) / 2
		switch guess(avg) {
		case -1:
			max = avg - 1
		case 1:
			min = avg + 1
		default:
			return avg
		}
	}

	return -1
}

func guess(n int) int {
	if n > target {
		return -1
	} else if n < target {
		return 1
	}
	return 0

}
