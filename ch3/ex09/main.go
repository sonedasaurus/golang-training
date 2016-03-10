package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

type Params struct {
	x, y, scale float64
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		params := Params{x: 0, y: 0, scale: 1} // default parameters
		if r.FormValue("x") != "" {
			params.x, _ = strconv.ParseFloat(r.FormValue("x"), 64)
		}
		if r.FormValue("y") != "" {
			params.y, _ = strconv.ParseFloat(r.FormValue("y"), 64)
		}
		if r.FormValue("scale") != "" {
			params.scale, _ = strconv.ParseFloat(r.FormValue("scale"), 64)
		}
		fractal(w, params)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

const (
	contrast = 15
	th       = 1
	size     = 1024
)

func fractal(w io.Writer, params Params) {
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	for py := 0; py < size; py++ {
		y := realize(py, params.scale) + params.y
		for px := 0; px < size; px++ {
			x := realize(px, params.scale) + params.x
			z := complex(x, y)

			count := solve(z)
			color := color.Gray{uint8(count) * contrast}

			img.Set(px, py, color)
		}
	}
	png.Encode(w, img)
}

func realize(d int, scale float64) float64 {
	width := 10 / scale
	return float64(d)/size*width - width/2
}

func solve(zn complex128) (ct uint64) {
	for cmplx.Abs(zn) > th {
		zn = zn - f(zn)/df(zn)
		ct++
	}

	return
}

func f(z complex128) complex128 {
	return cmplx.Pow(z, 4) - 1
}

func df(z complex128) complex128 {
	return 4 * cmplx.Pow(z, 3)
}
