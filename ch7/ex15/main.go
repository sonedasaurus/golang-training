package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./eval"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("式を入力してください")
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	expr, err := eval.Parse(s)
	if err != nil {
		fmt.Errorf("%s", err) // parse error
		os.Exit(1)
	}

	fmt.Println("変数の値を入力してください")
	var m eval.Env = map[eval.Var]float64{}
	slice := expr.Vars()
	for _, v := range slice {
		fmt.Printf("%s: ", v)
		var s string
		if sc.Scan() {
			s = sc.Text()
		}
		value, _ := strconv.ParseFloat(s, 64)
		m[eval.Var(v)] = value
	}
	fmt.Println(expr.Eval(m))
}
