/*

Exercise: Fibonacci closure

Let's have some fun with functions.

Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

*/

package functions

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	return nil //this line is just to eliminate the error
}

func Fibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// ./prog.go:8:1: missing return
