// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/daniruizcamacho/gopl.io/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		fahrenheit := tempconv.Fahrenheit(v)
		celsius := tempconv.Celsius(v)
		feets := tempconv.Feets(v)
		meters := tempconv.Meters(v)
		pounds := tempconv.Pounds(v)
		kilograms := tempconv.Kilograms(v)

		fmt.Printf("%s = %s, %s = %s\n",
			fahrenheit, tempconv.FToC(fahrenheit), celsius, tempconv.CToF(celsius))
		fmt.Printf("%s = %s, %s = %s\n",
			feets, tempconv.FToM(feets), meters, tempconv.MToF(meters))
		fmt.Printf("%s = %s, %s = %s\n",
			pounds, tempconv.PToK(pounds), kilograms, tempconv.KToP(kilograms))
	}
}

//!-
