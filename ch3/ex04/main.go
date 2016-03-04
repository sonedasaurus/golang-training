package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type Params struct {
	color, width, height string
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		params := Params{width: "600", height: "320", color: "white"} // default parameters
		if r.FormValue("width") != "" {
			params.width = r.FormValue("width")
		}
		if r.FormValue("height") != "" {
			params.height = r.FormValue("height")
		}
		if r.FormValue("color") != "" {
			params.color = r.FormValue("color")
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		eggbox(w, params)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func eggbox(out io.Writer, p Params) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
		"width='%s' height='%s'>", p.color, p.width, p.height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func selectColor(h float64) string {
	t := 255 * ((h / 0.2) - 1)
	s := ""
	if t > 0 {
		s = fmt.Sprintf("%02x", int(255-t))
		return "#ff" + s + s
	}
	s = fmt.Sprintf("%02x", int(255+t))
	return "#" + s + s + "ff"
}

func f(x, y float64) float64 {
	v := math.Pow(2.0, math.Sin(y)) * math.Pow(2.0, math.Sin(x)) / 12
	if math.IsNaN(v) {
		return 0
	}
	return v
}
