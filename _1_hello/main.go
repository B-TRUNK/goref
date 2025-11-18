package main

import (
	"fmt"

	"hello/morestrings"

	//Import Remote Packagers
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, Go!"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
