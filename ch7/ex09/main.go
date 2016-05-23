package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type Tracks struct {
	Tracks []*Track
	Keys   []string
}

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>Track List</h1>
<table>
<tr style='text-align: left'>
  <th><a href="?Title=true">Title</a></th>
  <th><a href="?Artist=true">Artist</a></th>
  <th><a href="?Album=true">Album</a></th>
  <th><a href="?Year=true">Year</a></th>
  <th><a href="?Length=true">Length</a></th>
</tr>
  {{range .Tracks}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func main() {
	var tracks = Tracks{Tracks: []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}, Keys: nil}
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("Title") != "" {
			tracks.Keys = append(tracks.Keys, "Title")
		}
		if r.FormValue("Artist") != "" {
			tracks.Keys = append(tracks.Keys, "Artist")
		}
		if r.FormValue("Album") != "" {
			tracks.Keys = append(tracks.Keys, "Album")
		}
		if r.FormValue("Year") != "" {
			tracks.Keys = append(tracks.Keys, "Year")
		}
		if r.FormValue("Length") != "" {
			tracks.Keys = append(tracks.Keys, "Length")
		}
		sort.Sort(customSort{tracks.Tracks, tracks.Keys, multiTier})
		if err := issueList.Execute(w, tracks); err != nil {
			log.Fatal(err)
		}
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func multiTier(x, y *Track, sortKeys []string) bool {
	for _, sortKey := range sortKeys {
		if sortKey == "Title" && x.Title != y.Title {
			return x.Title < y.Title
		}
		if sortKey == "Artist" && x.Artist != y.Artist {
			return x.Artist < y.Artist
		}
		if sortKey == "Album" && x.Album != y.Album {
			return x.Album < y.Album
		}
		if sortKey == "Year" && x.Year != y.Year {
			return x.Year < y.Year
		}
		if sortKey == "Length" && x.Length != y.Length {
			return x.Length < y.Length
		}
	}
	return false
}

type customSort struct {
	t        []*Track
	sortKeys []string
	less     func(x, y *Track, sortKeys []string) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j], x.sortKeys) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
