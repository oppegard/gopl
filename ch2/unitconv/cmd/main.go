package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch2/tempconv"
	"gopl.io/ch2/unitconv"
)

func main() {
	if len(os.Args) == 1 {
		reader := bufio.NewReader(os.Stdin)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
				os.Exit(1)
			}
			ConvertUnits(strings.TrimSuffix(line, "\n"))
		}
	} else {
		for _, arg := range os.Args[1:] {
			ConvertUnits(arg)
		}
	}
}
func ConvertUnits(arg string) {
	val, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(val)
	c := tempconv.Celsius(val)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
	m := unitconv.Meter(val)
	ft := unitconv.Foot(val)
	fmt.Printf("%s = %s, %s = %s\n",
		m, unitconv.MToF(m), ft, unitconv.FToM(ft))
	k := unitconv.Kilogram(val)
	p := unitconv.Pound(val)
	fmt.Printf("%s = %s, %s = %s\n",
		k, unitconv.KToP(k), p, unitconv.PToK(p))
}
