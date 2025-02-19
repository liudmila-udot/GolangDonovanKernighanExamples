// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/url"
	"os"
	"strconv"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main

var palette = []color.Color{color.Black,
	// green
	color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 255,
	},
	// red
	color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	},
	// blue
	color.RGBA{
		R: 0,
		G: 0,
		B: 255,
		A: 255,
	},
	// yellow
	color.RGBA{
		R: 255,
		G: 255,
		B: 0,
		A: 255,
	},
}

const (
	backgroundColourIndex = 0 // first color in palette
	lineColourIndex       = 1 // next color in palette
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w, r)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	//lissajous(os.Stdout)
}

func lissajous(out io.Writer, r *http.Request) {
	cycles := 5
	const (
		//cycles  = 20    // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	fmt.Println(r.URL.RawQuery)
	m, _ := url.ParseQuery(r.URL.RawQuery)
	fmt.Println(m)
	if m["cycles"] != nil {
		var err error
		cycles, err = strconv.Atoi(m["cycles"][0])
		check(err)
	}

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colourIndex := uint8(rand.Intn(5) + 1)
			//fmt.Printf("CoulourIndex: %d", colourIndex)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				colourIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//!-main
