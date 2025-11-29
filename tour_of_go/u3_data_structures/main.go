package main

import (
	"ds/arrays"
	"ds/maps"
	"ds/pointers"
	"ds/strct"
	"fmt"
)

func main() {

	fmt.Println("This Unit is All About Data-Structures")

	fmt.Println("Pointers :")
	pointers.PNT()

	strct.STRUCTURES()

	fmt.Println("Arrays :")
	arrays.Array()
	arrays.Slices()
	arrays.Slices_as_refs()
	arrays.Slices_as_literals()
	arrays.Slices_as_defaults()
	arrays.Slices_len_cap()
	arrays.Slices_nil()
	arrays.Slices_with_make()
	arrays.Slices_of_slices()
	arrays.Slices_appending()

	fmt.Println("Ranges: ")
	arrays.Range()
	arrays.Range2()
	//arrays.Pic()

	fmt.Println("Maps :")
	maps.Maps()

}
