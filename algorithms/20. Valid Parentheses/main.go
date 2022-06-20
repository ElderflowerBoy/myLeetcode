package main

func main() {
	str := ")"
	println(isValid(str))
	str = "()[]{}"
	println(isValid(str))
	str = "(]"
	println(isValid(str))
	str = "()"
	println(isValid(str))

}

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}

	var stack []string

	for i := 0; i < len(s); i++ {
		if isOpen(string(s[i])) {
			stack = append(stack, string(s[i]))
		} else {
			n := len(stack) - 1
			if n < 0 || getClosing(stack[n]) != string(s[i]) {
				return false
			}
			stack = stack[:n]
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}

func isOpen(s string) bool {
	switch s {
	case "(",
		"[",
		"{":
		return true
	default:
		return false

	}
}
func getClosing(s string) string {
	switch s {
	case "(":
		return ")"
	case "[":
		return "]"
	case "{":
		return "}"

	}
	return ""
}
