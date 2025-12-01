/*

Exercise: Maps

Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.

*/

package maps

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func Maps_Ex() {
	wc.Test(WordCount)
}

/*

FAIL
 f("I am learning Go!") =
  map[string]int{"x":1}
 want:
  map[string]int{"Go!":1, "I":1, "am":1, "learning":1}

*/
