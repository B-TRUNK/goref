/*

For is Go's "while"

At that point you can drop the semicolons: C's while is spelled for in Go.

*/

package forloop

import "fmt"

func For_As_While() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// 1024
