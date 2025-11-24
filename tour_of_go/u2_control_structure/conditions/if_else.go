/*

If and else

Variables declared inside an if short statement are also available inside any of the else blocks.

(Both calls to pow return their results before the call to fmt.Println in main begins.)

*/

package conditions

import (
	"fmt"
	"math"
)

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func IF_ELSE() {
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}
