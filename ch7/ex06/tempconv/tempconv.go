package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + 273.5) }
func KToC(k Kelvin) Celsius     { return Celsius(k - 273.5) }
func KToF(k Kelvin) Fahrenheit  { return Fahrenheit((k-273.5)*9.0/5.0 + 32.0) }
func FToK(f Fahrenheit) Kelvin  { return Kelvin(((f - 32.0) * 5.0 / 9.0) + 273.5) }

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

type celsiusFlag struct{ Celsius }
type kelvinFlag struct{ Kelvin }
type fahrenheitFlag struct{ Fahrenheit }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func (f *kelvinFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Kelvin = CToK(Celsius(value))
		return nil
	case "F", "°F":
		f.Kelvin = FToK(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Kelvin = Kelvin(value)
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func (f *fahrenheitFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Fahrenheit = CToF(Celsius(value))
		return nil
	case "F", "°F":
		f.Fahrenheit = Fahrenheit(value)
		return nil
	case "K", "°K":
		f.Fahrenheit = KToF(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func KelvinFlag(name string, value Kelvin, usage string) *Kelvin {
	f := kelvinFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Kelvin
}

func FahrenheitFlag(name string, value Fahrenheit, usage string) *Fahrenheit {
	f := fahrenheitFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Fahrenheit
}
