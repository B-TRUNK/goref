/*

Map literals

Map literals are like struct literals, but the keys are required.

*/

package maps

import "fmt"

type Pack struct {
	Lat, Long float64
}

var map_literal = map[string]Pack{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func Map_Literal() {
	fmt.Println(map_literal)
}
