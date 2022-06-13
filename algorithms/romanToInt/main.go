package main

import "fmt"

var digitMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func main() {
	str := "LXVII"
	fmt.Println(romanToInt(str))
}

func romanToInt(s string) int {
	var number int

	for i := 0; i < len(s); i++ {

		var curr = digitMap[string(s[i])]
		var next int
		if i < len(s)-1 {
			next = digitMap[string(s[i+1])]
		}

		if next > 0 && curr < next && ((next == 5 && curr == 1) || next%10 == 0) {
			number += next - curr
			i++
			continue
		}

		number += curr
	}
	return number
}
