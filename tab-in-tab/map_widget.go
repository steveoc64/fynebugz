package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
)

// MapWidget is a complete map viewer widget
// ... or will be when it grows up
type MapWidget struct {
	size     fyne.Size
	position fyne.Position
	hidden   bool
	mapData string
	other *MapWidget
}

func newMapWidget(mapData string, other *MapWidget) *MapWidget {
	mw := &MapWidget{
		mapData: mapData,
		other: other,
	}
	mw.Resize(mw.MinSize())
	return mw
}

// Size returns the current size of the mapWidget
func (mw *MapWidget) Size() fyne.Size {
	return mw.size
}

// Resize resizes the mapWidget
func (mw *MapWidget) Resize(size fyne.Size) {
	println("mw resize to", size.Width, size.Height)
	mw.size = size
	widget.Renderer(mw).Layout(mw.size)
	canvas.Refresh(mw)
}

// Position returns the current position of the mapWidget
func (mw *MapWidget) Position() fyne.Position {
	return mw.position
}

// Move orders the mapWidget to be moved
func (mw *MapWidget) Move(pos fyne.Position) {
	mw.position = pos
	widget.Renderer(mw).Layout(mw.size)
}

// MinSize returns the minSize of the mapWitdget
func (mw *MapWidget) MinSize() fyne.Size {
	return widget.Renderer(mw).MinSize()
}

// Visible returns whether the mapWidget is visible or not
func (mw *MapWidget) Visible() bool {
	return !mw.hidden
}

// Show sets the mapWidget to be visible
func (mw *MapWidget) Show() {
	println("widget show")
	if mw.other != nil {
		mw.other.Hide()
	}
	mw.hidden = false
}

// Hide sets the mapWidget to be not visible
func (mw *MapWidget) Hide() {
	println("widget hide")
	if mw.other != nil {
		mw.other.Show()
	}
	mw.hidden = true
}

// ApplyTheme applies the theme to the mapWidget
func (mw *MapWidget) ApplyTheme() {
	widget.Renderer(mw).ApplyTheme()
}

// CreateRenderer builds a new renderer
func (mw *MapWidget) CreateRenderer() fyne.WidgetRenderer {
	return newMapRender(mw)
}