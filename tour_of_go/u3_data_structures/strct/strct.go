/*

Structs

A struct is a collection of fields.

*/

package strct

import "fmt"

type Vertex struct {
	X int
	Y int
}

// Struct Literals: 3rd Section
var (
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X: 1} // Y:0 is implicit
	v3 = Vertex{}     // X:0 and Y:0
)

func STRUCTURES() {
	fmt.Println(Vertex{1, 2})

	//strct_fields
	v := Vertex{5, 6}
	v.Y = 4
	fmt.Println("X = ", v.X)
	fmt.Println("Y = ", v.Y)

	/*

		Pointers to structs

		Struct fields can be accessed through a struct pointer.

		To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

	*/
	p := &v
	p.X = 1e9
	fmt.Println("Pointer to Struct: ", v)

	/*

		Struct Literals

		A struct literal denotes a newly allocated struct value by listing the values of its fields.

		You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

		The special prefix & returns a pointer to the struct value.

	*/

	p = &Vertex{1, 2} // has type *Vertex
	fmt.Println("Struct Literals", v1, p, v2, v3)

}
