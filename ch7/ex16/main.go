package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"./eval"
)

var calculate = template.Must(template.New("calculate").Parse(`
<html>
<head>
<title></title>
</head>
<body>
<form action="/" method="post">
  式を入力してください:<input type="text" name="expr">
  <input type="submit" value="決定">
</form>
<p>{{.}}</p>
</body>
</html>
`))

func main() {
	var anser float64
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("expr") != "" {
			expr, err := eval.Parse(r.FormValue("expr"))
			if err != nil {
				fmt.Errorf("%s", err) // parse error
				os.Exit(1)
			}
			anser = expr.Eval(nil)
		}
		if err := calculate.Execute(w, anser); err != nil {
			log.Fatal(err)
		}
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}
