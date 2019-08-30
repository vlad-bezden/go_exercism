//Package grains provides solution for "grains" exercise on Exercism
package grains

import "errors"

//Square calculate the number of grains of wheat on n square of chessboard
func Square(n int) (uint64, error) {
	if n <= 0 || 64 < n {
		return 0, errors.New("invalid input. Valid values 0-64")
	}

	return 1 << uint(n-1), nil
}

//Total return total number of grains
func Total() uint64 {
	return 1<<uint(64) - 1
}
