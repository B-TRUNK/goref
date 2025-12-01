/*
Variables
The var statement declares a list of variables; as in function argument lists,
 the type is last.
A var statement can be at package or function level. We see both in this example.
*/

package multivariables

import "fmt"

var c, python, java bool

func Display() {
	var i int
	fmt.Println(i, !c, python, java)
}

// 55
