// Package gocmdflags is the code of flags value in the go programming language
package gocmdflags

import (
	"flag"
	"fmt"
)

// Celsius the test type for cmd flags
type Celsius float64

// Fahrenheit the test type for cmd flags
type Fahrenheit float64

// Kelvin the test type for cmd flags
type Kelvin float64

const (
	absoluteZeroC Celsius = -273.15
	freezingC     Celsius = 0
	boilingC      Celsius = 100
)

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = ftoC(Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = kToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func cToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func ftoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func kToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// CelsiusFlag defines a Celius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C"
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
