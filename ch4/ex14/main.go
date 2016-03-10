package main

import (
	"log"
	"net/http"

	"./github"
)

//!+template
import "html/template"

var issueTempl = template.Must(template.New("issue").Parse(`
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
</table>
`))

var milestoneTempl = template.Must(template.New("issue").Parse(`
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>Title</th>
</tr>
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
  <td>{{.Description}}</td>
</tr>
</table>
`))

func main() {
	listResult, err := github.ListIssues()
	if err != nil {
		log.Fatal(err)
	}
	Milestone, err := github.Milestones()
	if err != nil {
		log.Fatal(err)
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		for _, item := range listResult {
			if err := issueTempl.Execute(w, item); err != nil {
				log.Fatal(err)
			}
		}
		for _, item := range Milestone {
			if err := milestoneTempl.Execute(w, item); err != nil {
				log.Fatal(err)
			}
		}
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}
