package gutil

import "unicode/utf8"

// Clamp returns min if v is less than min, max if v is greater than max, otherwise v.
// This function is used to clamp a value to a range.
// Based on CSS clamp function.
// https://developer.mozilla.org/en-US/docs/Web/CSS/clamp
// Example:
// 	Clamp(10, 0, 20) // returns 10
func Clamp(min, v, max int) int {
	if min > max {
		min, max = max, min
	}
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func HasNonASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > 0x7F {
			// Fast path: any byte > 0x7F implies non-ASCII UTF-8
			return true
		}
	}
	return false
}

func HasNonASCIIValidUTF8(s string) (nonASCII bool, okUTF8 bool) {
	if !utf8.ValidString(s) {
		return false, false
	}
	for i := 0; i < len(s); i++ {
		if s[i] > 0x7F {
			return true, true
		}
	}
	return false, true
}
