package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y1 := float64(py*2)/height*(ymax-ymin) + ymin
		y2 := float64(py*2-1)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x1 := float64(px*2)/width*(xmax-xmin) + xmin
			x2 := float64(px*2-1)/width*(xmax-xmin) + xmin
			z1 := complex(x1, y1)
			z2 := complex(x1, y2)
			z3 := complex(x2, y1)
			z4 := complex(x2, y2)
			grayScale := (mandelbrot(z1) + mandelbrot(z2) + mandelbrot(z3) + mandelbrot(z4)) / 4
			img.Set(px, py, color.Gray{grayScale})
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) uint8 {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return 255 - (contrast * n)
		}
	}
	return 255
}
