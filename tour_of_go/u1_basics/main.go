package main

import (
	"bsq/functions"
	"bsq/packages"
	"bsq/testing"
	"bsq/types"
	"bsq/vars"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
	/*
		Welcome to the playground!
		The time is 2025-12-01 20:37:25.443156944 +0200 EET m=+0.000145169
	*/

	packages.Pkgs()
	packages.Imports()
	packages.Exported_Names()

	functions.Func()
	functions.Func_Continued()
	functions.Swap()
	functions.Func_Named_Ret_Values()

	vars.Vars()
	vars.Var_Inits()
	vars.Vars_Short_Declaration()

	types.Types()
	types.Zero_Vals()
	types.Type_Conversions()
	types.Inference()
	types.Constants()
	types.Numeric_Constants()

	fmt.Println(testing.ReverseRunes("Hello, World!"))

}
