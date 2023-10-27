// Package provides various math algorithms.
package algo

func Euclid(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	if a < 0 {
		a *= -1
	}

	if b < 0 {
		b *= -1
	}

	if a == b {
		return a
	}

	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}

	return a
}
