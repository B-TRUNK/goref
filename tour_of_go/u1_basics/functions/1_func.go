/*


Functions

A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.

(For more about why types look the way they do, see the article on Go's declaration syntax.)

*/

package functions

import "fmt"

func add(x, y int) int {
	return x + y
}

func Func() {
	fmt.Println(add(42, 13))
}

// 55
