package main

import (
	"control/conditions"
	"control/forloop"
	"fmt"
)

func main() {

	fmt.Println("This Module is all About")

	forloop.ForLoop()
	forloop.ForLoop_Con()
	forloop.For_As_While()

	conditions.SQRT()
	conditions.IF_SHORT()
	conditions.IF_ELSE()

	conditions.SWITCH_CASE()
	conditions.SWITCH_VAL_ORDER()

	conditions.DEFER()
	conditions.DEFER_STACKING()

}
