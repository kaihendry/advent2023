// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package advent

import "testing"

func TestString(t *testing.T) {
	for _, c := range []struct {
		in   string
		want int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	} {
		got := Trebuchet(c.in)
		if got != c.want {
			t.Errorf("Trebuchet(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
