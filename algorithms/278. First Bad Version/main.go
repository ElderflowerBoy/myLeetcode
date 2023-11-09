package main

const badVersion = 4

func main() {
	firstBadVersion(7)
}

func firstBadVersion(version int) int {
	min := 1
	max := version

	for min <= max {
		avg := (min + max) / 2

		if isBadVersion(avg) {
			max = avg - 1
		} else {
			min = avg + 1
		}
	}

	return min
}

func isBadVersion(version int) bool {
	return version >= badVersion
}
