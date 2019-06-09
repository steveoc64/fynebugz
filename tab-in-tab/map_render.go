package main

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"github.com/davecgh/go-spew/spew"
	"github.com/llgcode/draw2d/draw2dimg"
)

type mapRender struct {
	render  *canvas.Raster
	mw      *MapWidget
	objects []fyne.CanvasObject
	img     *image.RGBA
	X       int
	Y       int
}

func newMapRender(mw *MapWidget) *mapRender {
	r := &mapRender{mw: mw, X: 6, Y: 4}
	render := canvas.NewRaster(r.getImage)
	//render := canvas.NewRasterWithPixels(r.paint2)
	r.render = render
	r.objects = []fyne.CanvasObject{render}
	//r.ApplyTheme()
	return r
}

// ApplyTheme applies the theme
func (r *mapRender) ApplyTheme() {
	// noop
}

// BackgroundColor returns the background color for our map
func (r *mapRender) BackgroundColor() color.Color {
	return color.RGBA{183, 211, 123, 1}
}

// Destroy removes any resources we have on this renderer
func (r *mapRender) Destroy() {
	// noop
}

// Layout does .. the layout ?
func (r *mapRender) Layout(size fyne.Size) {
	println("renderer layout", size.Width, size.Height)
	r.render.Resize(size)
}

// MinSize returns the minimum size for this renderer
func (r *mapRender) MinSize() fyne.Size {
	return fyne.Size{
		Width:  int(r.X * 64),
		Height: int(r.Y * 64),
	}
}

// Objects returns the slice of objects that we own
func (r *mapRender) Objects() []fyne.CanvasObject {
	return r.objects
}

// Refresh paints the map
func (r *mapRender) Refresh() {
	println("maprender refresh")
	canvas.Refresh(r.mw)
}

func (r *mapRender) getImage(w, h int) image.Image {
	println("been asked to getImage and hidden =", r.mw.hidden, w, h)
	if r.img != nil {
		spew.Dump(r.img.Bounds())
	}
	if r.img == nil || r.img.Bounds().Size().X != w || r.img.Bounds().Size().Y != h {
		r.img = r.generateImage(w, h)
	}
	// this wont exactly fix the problem, but it will hide it !
	if false {
		if r.mw.hidden {
			return &image.RGBA{}
		}
	}
	return r.img
}

func (r *mapRender) generateImage(w, h int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	if w == 0 || h == 0 {
		return img
	}
	dx := float64(w / int(r.X))
	dy := float64(h / int(r.Y))
	mx := int(r.X)
	my := int(r.Y)
	gc := draw2dimg.NewGraphicContext(img)

	// grid overlays
	i := 0
	for y := 0; y < my; y++ {
		for x := 0; x < mx; x++ {
			c := color.RGBA{uint8(rand.Intn(64)), 200, 50, uint8(rand.Intn(16))}
			draw.Draw(img,
				image.Rectangle{image.Point{x * int(dx), y * int(dy)},
					image.Point{(x + 1) * int(dx), (y + 1) * int(dy)}},
				&image.Uniform{c},
				image.Point{0, 0},
				draw.Src)
			i++
		}
	}

	// grid lines - vertical
	gc.SetStrokeColor(color.RGBA{32, 32, 32, 128})
	gc.SetLineWidth(1)
	for x := 0; x < mx; x++ {
		gc.BeginPath()
		gc.MoveTo(float64(x)*dx, 0.0)
		gc.LineTo(float64(x)*dx, float64(h))
		gc.Close()
		gc.FillStroke()
	}
	// grid lines - horizontal
	for y := 0; y < my; y++ {
		gc.BeginPath()
		gc.MoveTo(0.0, float64(y)*dy)
		gc.LineTo(float64(w), float64(y)*dy)
		gc.Close()
		gc.FillStroke()
	}

	return img
}
