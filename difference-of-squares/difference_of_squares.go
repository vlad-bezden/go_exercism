//Package diffsquares provides solution for 'Difference Of Squares" exersice on Exercism
package diffsquares

//SquareOfSum finds square of the sum
//The square of the sum of the first ten natural numbers is
//(1 + 2 + ... + 10)² = 55² = 3025.
func SquareOfSum(n int) (sum int) {
	sum = n * (n + 1) / 2
	return sum * sum
}

//SumOfSquares finds the sum of squares
//The sum of the squares of the first ten natural numbers is
//1² + 2² + ... + 10² = 385.
func SumOfSquares(n int) (sum int) {
	return n * (n + 1) * (2*n + 1) / 6
}

//Difference finds the difference between square of sum
//and sum of squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
