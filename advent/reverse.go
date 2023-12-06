package advent

import (
	"strings"
)

// returns the numbers in the string
func Trebuchet(s string) int {
	var output []int
	var letters string
	for _, r := range s {
		if r >= '0' && r <= '9' {
			output = append(output, int(r-'0'))
			letters = ""
		} else {
			// add letter to string
			letters += string(r)
			// log.Println("adding letter", letters, string(r))
		}
		// check if string is a number
		numMap := map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
			"six":   6,
			"seven": 7,
			"eight": 8,
			"nine":  9,
		}

		for key, value := range numMap {
			if strings.Contains(letters, key) {
				output = append(output, value)
				letters = ""
			}
		}
	}

	// first item of slice is the first digit
	first := output[0]
	// last item of slice is the last digit
	last := output[len(output)-1]

	// return first as first digit and last as last digit
	return first*10 + last
}
