// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
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
func main() {
	fmt.Println("\n1: Title 2: Artist\n")
	sortKeys := []string{"Title", "Artist"}
	sort.Sort(customSort{tracks, sortKeys, multiTier})
	printTracks(tracks)

	fmt.Println("\n1: Title 2: Year\n")
	sortKeys = []string{"Title", "Year"}
	sort.Sort(customSort{tracks, sortKeys, multiTier})
	printTracks(tracks)

	fmt.Println("\n1: Album 2: Year\n")
	sortKeys = []string{"Album", "Year"}
	sort.Sort(customSort{tracks, sortKeys, multiTier})
	printTracks(tracks)
}

/*
//!+customout
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go          Moby            Moby               1992  3m37s
Go          Delilah         From the Roots Up  2012  3m38s
Go Ahead    Alicia Keys     As I Am            2007  4m36s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
//!-customout
*/

type customSort struct {
	t        []*Track
	sortKeys []string
	less     func(x, y *Track, sortKeys []string) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j], x.sortKeys) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
