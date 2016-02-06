package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

type Params struct {
	cycles, res, size float64
	nframes, delay    int
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	handler := func(w http.ResponseWriter, r *http.Request) {
		params := Params{cycles: 5, res: 0.001, size: 100, nframes: 64, delay: 8} // default parameters
		if r.FormValue("cycles") != "" {
			params.cycles, _ = strconv.ParseFloat(r.FormValue("cycles"), 64)
		}
		if r.FormValue("res") != "" {
			params.res, _ = strconv.ParseFloat(r.FormValue("res"), 64)
		}
		if r.FormValue("size") != "" {
			params.size, _ = strconv.ParseFloat(r.FormValue("size"), 64)
		}
		if r.FormValue("nframes") != "" {
			params.nframes, _ = strconv.Atoi(r.FormValue("nframes"))
		}
		if r.FormValue("delay") != "" {
			params.delay, _ = strconv.Atoi(r.FormValue("delay"))
		}
		lissajous(w, params)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func lissajous(out io.Writer, p Params) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: p.nframes}
	phase := 0.0 // phase difference
	for i := 0; i < p.nframes; i++ {
		rect := image.Rect(0, 0, 2*int(p.size)+1, 2*int(p.size)+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < p.cycles*2*math.Pi; t += p.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(int(p.size+(x*p.size+0.5)), int(p.size+(y*p.size+0.5)),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, p.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
