package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout // modified during testing

func main() {
	echo(os.Args[0:])
}

func echo(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Fprint(out, s)
}
