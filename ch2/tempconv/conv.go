// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 275.15) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 275.15) }

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(k*9/5 - 459.67) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return Kelvin((f + 459.67) * 5 / 9) }

// MToF converts a Feets distance to Meters.
func MToF(m Meters) Feets { return Feets(m * 3.281) }

// FToM converts a Meters distance to Feets.
func FToM(f Feets) Meters { return Meters(f / 3.281) }

// PToK converts a Pounds weigh to Kilograms.
func PToK(p Pounds) Kilograms { return Kilograms(p / 2.205) }

// KToP converts a Kilograms weigh to Pounds.
func KToP(k Kilograms) Pounds { return Pounds(k * 2.205) }

//!-
