package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"./github"
)

const searchTempl = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

const listTempl = `---------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var searchReport = template.Must(template.New("search").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(searchTempl))

var listReport = template.Must(template.New("list").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(listTempl))

func main() {
	if os.Args[1] == "create" {
		err := github.CreateIssues(os.Args[2])
		if err != nil {
			fmt.Println(err)
		}
	}
	if os.Args[1] == "list" {
		result, err := github.ListIssues()
		if err != nil {
			log.Fatal(err)
		}
		for _, item := range result {
			if err := listReport.Execute(os.Stdout, item); err != nil {
				log.Fatal(err)
			}
		}
	}
	if os.Args[1] == "search" {
		result, err := github.SearchIssues(os.Args[1:])
		if err != nil {
			log.Fatal(err)
		}
		if err := searchReport.Execute(os.Stdout, result); err != nil {
			log.Fatal(err)
		}
	}
	if os.Args[1] == "edit" {
		err := github.EditIssues(os.Args[2], os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
	}
}
