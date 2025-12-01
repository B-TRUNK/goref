/*

Multiple results

A function can return any number of results.

The swap function returns two strings.

*/

package functions

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func Swap() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

// world hello
