package main

import (
	"fmt"

	"gopl.io/ch2/tempconv"
)

func main() {
	fmt.Printf("Brrr! %v\n", tempconv.AbsoluteZeroC)
	k := tempconv.Kelvin(50)
	c := tempconv.Celsius(0)
	f := tempconv.Fahrenheit(0)

	fmt.Printf("%v: %v\n", k, tempconv.KToC(k))
	fmt.Printf("%v: %v\n", k, tempconv.KToF(k))

	fmt.Printf("%v: %v\n", c, tempconv.CToK(c))
	fmt.Printf("%v: %v\n", f, tempconv.FToK(f))
}
