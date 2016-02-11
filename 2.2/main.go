package main

import (
	"fmt"
	"os"
	"strconv"

	"../2.1/tempconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		arg := ""
		fmt.Scan(&arg)
		args = append(args, arg)
	}
	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
