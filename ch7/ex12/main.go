package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/register", db.register)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var dbList = template.Must(template.New("dblist").Parse(`
<h1>DB List</h1>
<table>
<tr style='text-align: left'>
  <th>Item</th>
  <th>Price</th>
</tr>
  {{range $index, $element := .}}
<tr>
  <td>{{$index}}</td>
  <td>{{$element}}</td>
</tr>
{{end}}
</table>
`))

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := dbList.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	p := req.URL.Query().Get("price")
	price, _ := strconv.ParseFloat(p, 32)
	if oldPrice, ok := db[item]; ok {
		db[item] = dollars(price)
		fmt.Fprintf(w, "Changed %d price %f from %s\n", item, price, oldPrice)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "%q is deleted\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) register(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	p := req.URL.Query().Get("price")
	price, _ := strconv.ParseFloat(p, 32)
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest) // 401
		fmt.Fprintf(w, " %q is already exists\n", item)
	} else {
		db[item] = dollars(price)
		fmt.Fprintf(w, " %q is registered\n", item)
	}
}
