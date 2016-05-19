package main

import (
	"flag"
	"fmt"

	"./tempconv"
)

var ctemp = tempconv.CelsiusFlag("c", 20.0, "the temperature in celsius")
var ktemp = tempconv.KelvinFlag("k", 293.05, "the temperature in kelvin")
var ftemp = tempconv.FahrenheitFlag("f", 68.0, "the temperature in fahrenheit")

func main() {
	flag.Parse()
	fmt.Println(*ctemp)
	fmt.Println(*ktemp)
	fmt.Println(*ftemp)
}
