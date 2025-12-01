/*

Imports

This code groups the imports into a parenthesized, "factored" import statement.

You can also write multiple import statements, like:

import "fmt"
import "math"

But it is good style to use the factored import statement.

*/

package packages

import (
	"fmt"
	"math"
)

func Imports() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

// Now you have 2.6457513110645907 problems.
