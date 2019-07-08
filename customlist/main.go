package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	println("custom list exmaple")

	app := app.New()
	w := app.NewWindow("custom list rendering")
	w.SetContent(newCustomListUI())
	w.CenterOnScreen()
	w.ShowAndRun()
}

// our data model - just a list of things
var items []string
var itemID = 1

func newCustomListUI() *fyne.Container {
	header := widget.NewHBox(widget.NewLabel("The Header"))
	footer := widget.NewHBox(widget.NewLabel("The Footer"))

	entry := widget.NewButton("Add another item to the data model", func() {
		items = append(items, fmt.Sprintf("This is item number %d", len(items)+1))
		println("there are now", len(items), "things in the datamodel")
	})
	clearBtn := widget.NewButton("Clear all the items", func() {
		items = []string{}
		println("just cleared all the items")
	})
	l := layout.NewBorderLayout(header, footer, nil, nil)
	tabContainer := widget.NewTabContainer(
		widget.NewTabItem("Add Items", entry),
		widget.NewTabItem("Show Items", newCustomList()),
		widget.NewTabItem("Clear Items", clearBtn),
	)
	tabContainer.SetTabLocation(widget.TabLocationLeading)
	return fyne.NewContainerWithLayout(l, header, footer, tabContainer)
}

// customList is an example of a Thin Wrapper around a standard widget - in this case
// a Box widget, as it embeds a box, and only overrides the Show() method to customise
// its contents when made visible.
type customList struct {
	widget.Box
	btn        *widget.Button
	firstLabel *widget.Label
}

func newCustomList() *customList {
	vbox := widget.NewVBox()
	l := &customList{
		Box: *vbox,
		btn: widget.NewButtonWithIcon("Do we have enough items yet", theme.CheckButtonIcon(), func() {
			println("We have", len(items), "items in the datamodel")
		}),
		firstLabel: widget.NewLabel("Here is a label to start with"),
	}
	l.Append(l.btn)
	l.Append(l.firstLabel)
	return l
}

func (l *customList) Show() {
	println("SHOW the list - so the tab just become visible")
	// update the button
	icon := theme.CheckButtonIcon()
	if len(items) > 3 {
		// we have 3 or more, so set the checked box
		icon = theme.CheckButtonCheckedIcon()
	}
	println("setting icon in the app::Show() to", icon.Name())
	l.btn.SetIcon(nil)
	l.btn.SetIcon(icon)
	l.btn.SetText(fmt.Sprintf("Is %d items enough yet ?", len(items)))
	l.btn.Show()

	// update the static label
	l.firstLabel.SetText(fmt.Sprintf("This label has been updated, and updates immediately"))

	// truncate the contents to 2 so we can append all the items onto the slice
	// NOTE - comment this l.Children line out, and you get the following behaviour :
	// - each Show() will paint the items that were added previously
	// - the new items take up extra space, but dont render in this call
	// - but they will render in the next call

	// If you un-comment this line, then you get this behaviour
	// - each Show() will take up extra space, because the new items layout
	// - but the new items will never paint

	l.Children = l.Children[:2]

	// add new contents
	for _, v := range items {
		fmt.Printf("appending a new label [%d]  with contents '%s'\n", itemID, v)
		l.Append(widget.NewLabel(fmt.Sprintf("[%d] %s", itemID, v)))
		itemID++
	}
	// paint it all
	l.Box.Show()
	widget.Refresh(l)
}
