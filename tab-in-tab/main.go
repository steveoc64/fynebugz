package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"os"
)

func main() {
	println("tab in tab example")
	println("Run with no args - uses stock label widgets, works OK")
	println("Run with -custom - uses custom label widgets, not OK")

	app := app.New()
	w := app.NewWindow("tab in tab bug")
	if len(os.Args) > 1 && os.Args[1] == "-custom" {
		w.SetContent(newTabInTabUICustom())
	} else {
		w.SetContent(newTabInTabUI())
	}
	w.CenterOnScreen()
	w.ShowAndRun()
}

func newTabInTabUI() *fyne.Container {
	header := widget.NewHBox(newTapLabel("The Header"))
	footer := widget.NewHBox(newTapLabel("The Footer"))
	tabPanel := []fyne.CanvasObject{
		widget.NewVBox(widget.NewLabel("zeroth first line"), widget.NewLabel("zeroth second line")),
		widget.NewVBox(widget.NewLabel("oneth first line"), widget.NewLabel("oneth second line")),
		widget.NewVBox(widget.NewLabel("tooth first line"), widget.NewLabel("tooth second line")),
		widget.NewVBox(widget.NewLabel("threeth first line"), widget.NewLabel("threeth second line")),
	}

	tabTab := widget.NewTabContainer(
		widget.NewTabItem("H1", tabPanel[2]),
		widget.NewTabItem("H2", widget.NewLabel("another label here")),
	)

	l := layout.NewBorderLayout(header, footer, nil, nil)
	tabContainer := widget.NewTabContainer(
		widget.NewTabItem("Tab 0", tabPanel[0]),
		widget.NewTabItem("Tab 1", tabPanel[1]),
		widget.NewTabItem("Tab Tab", tabTab),
		//widget.NewTabItem("Tab Tab", tabPanel[2]),   // still get the same issue without nested tabs too
		widget.NewTabItem("Tab 3", tabPanel[3]),
	)
	tabContainer.SetTabLocation(widget.TabLocationLeading)
	return fyne.NewContainerWithLayout(l, header, footer, tabContainer)
}


func newTabInTabUICustom() *fyne.Container {
	header := widget.NewHBox(newTapLabel("The Header"))
	footer := widget.NewHBox(newTapLabel("The Footer"))
	tabPanel := []fyne.CanvasObject{
		widget.NewVBox(newTapLabel("zeroth first line"), newTapLabel("zeroth second line")),
		widget.NewVBox(newTapLabel("oneth first line"), newTapLabel("oneth second line")),
		widget.NewVBox(newTapLabel("tooth first line"), newTapLabel("tooth second line")),
		widget.NewVBox(newTapLabel("threeth first line"), newTapLabel("threeth second line")),
	}

	tabTab := widget.NewTabContainer(
		widget.NewTabItem("H1", tabPanel[2]),
		widget.NewTabItem("H2", newTapLabel("another label here")),
		)

	l := layout.NewBorderLayout(header, footer, nil, nil)
	tabContainer := widget.NewTabContainer(
		widget.NewTabItem("Tab 0", tabPanel[0]),
		widget.NewTabItem("Tab 1", tabPanel[1]),
		widget.NewTabItem("Tab Tab", tabTab),
		widget.NewTabItem("Tab 3", tabPanel[3]),
	)
	tabContainer.SetTabLocation(widget.TabLocationLeading)
	return fyne.NewContainerWithLayout(l, header, footer, tabContainer)
}

type TappableLabel struct {
	widget.Label
}

func (t *TappableLabel) Tapped(event *fyne.PointEvent) {
	println("tapped label", t.Text)
}

func (t *TappableLabel) TappedSecondary(event *fyne.PointEvent) {
	println("secondary tap on", t.Text)
}

func newTapLabel(text string) *TappableLabel {
	return &TappableLabel{
		widget.Label{
			Text: text,
		},
	}
}
