package main

import "fmt"

func main() {
	var strings = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strings))
}

func longestCommonPrefix(strs []string) string {
	var prefix string
	var minLength = len(strs[0])
	for i := range strs {
		if len(strs[i]) < minLength {
			minLength = len(strs[i])
		}
	}

	for i := 0; i < minLength; i++ {
		prefix += string(strs[0][i])
		for _, str := range strs {
			if prefix[i] != str[i] {
				return prefix[:len(prefix)-1]
			}
		}
	}

	return prefix
}
