package advent

// returns the numbers in the string
func Trebuchet(s string) int {
	r := []rune(s)
	var output []int
	for i := 0; i < len(r); i++ {
		// if r[i] is a digit, add it to the slice
		if r[i] >= '0' && r[i] <= '9' {
			output = append(output, int(r[i]))
		}
	}

	// first item of slice is the first digit
	first := output[0] - '0'
	// last item of slice is the last digit
	last := output[len(output)-1] - '0'

	// return first as first digit and last as last digit
	return first*10 + last
}
