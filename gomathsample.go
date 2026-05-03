// This module contains simple math operations
package gomathsample

import "golang.org/x/exp/constraints" 

type Number interface{
	constraints.Integer | constraints.Float
}
// A function to add to numbers and return the results
// You can learn more about this operation at https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a T,b T) T{
	return a+b
}
