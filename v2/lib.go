// Package addition provides a function for adding two integers
package addition

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

// Add takes two integers and adds them together.
//
// It has two parameters: both parameters are integers, Add returns the sum of the
// two integer parameters
// [Addition](https://www.mathsisfun.com/numbers/addition.html)
func Add[T Number](a, b T) T {
	return a + b
}
