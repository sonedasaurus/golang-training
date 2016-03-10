package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"./xkcd"
)

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	Safe_title string
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

const comicTempl = `---------------------------------------
Number: {{.Num}}
Title:   {{.Title}}
Transcript:  {{.Transcript}}
`

var comicReport = template.Must(template.New("comic").
	Parse(comicTempl))

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./ex12 {search term}")
		os.Exit(1)
	}

	for i := 1; i < 50; i++ {
		filename := "./json/" + strconv.Itoa(i) + ".json"

		// file is Exist
		if _, err := os.Stat(filename); err == nil {
			file, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Fatal(err)
			}
			var comic Comic
			json.Unmarshal(file, &comic)

			// Report
			if strings.Contains(comic.Title, os.Args[1]) {
				if err := comicReport.Execute(os.Stdout, comic); err != nil {
					log.Fatal(err)
				}
			}
			continue
		}

		// Get Request
		result, err := xkcd.Get(i)
		if err != nil {
			log.Fatal(err)
		}

		// Report
		if strings.Contains(result.Title, os.Args[1]) {
			if err := comicReport.Execute(os.Stdout, result); err != nil {
				log.Fatal(err)
			}
		}

		// save file
		data, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		ioutil.WriteFile(filename, data, os.ModePerm)
	}
}
