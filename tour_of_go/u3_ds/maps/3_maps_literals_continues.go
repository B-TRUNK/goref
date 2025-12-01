/*

Map literals continued

If the top-level type is just a type name, you can omit it from the elements of the literal.

*/

package maps

import "fmt"

type Pack2 struct {
	Lat, Long float64
}

var map_literal2 = map[string]Pack2{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func Map_Literal2() {
	fmt.Println(map_literal2)
}

// map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
