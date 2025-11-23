/*

For continued

The init and post statements are optional.

*/

package forloop

import "fmt"

func ForLoop_Con() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
