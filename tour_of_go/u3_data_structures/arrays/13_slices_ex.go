/*

Exercise: Slices

Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)

*/

package arrays

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	return nil
}

func Slices_final_ex() {
	pic.Show(Pic)
}

/*

go: downloading golang.org/x/tour v0.0.0-20201207214521-004403599411
# example
./prog.go:6:1: missing return

*/
