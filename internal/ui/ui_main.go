package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func MainWindow() {
	application := app.NewWithID("Git")
	window := application.NewWindow("Git GUI")
	defer window.Close()
	window.SetTitle("Git")
	window.CenterOnScreen()
	var content = widget.SplitContainer{}
	// unstagedBox := widget.NewVBox(widget.NewLabel("Unstaged"))
	// stagedBox := widget.NewVBox(widget.NewLabel("Staged"))
	// commitedBox := widget.NewVBox(widget.NewLabel("Committed"))
	resultBox := widget.NewVBox(widget.NewLabel("Results"))
	// cont := fyne.NewContainer(nil)
	statusPane := fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(3), unstagedBox(window, resultBox), stagedBox(window, resultBox), commitedBox(window, resultBox))
	// actionPane := widget.NewVScrollContainer(widget.NewLabel("Action"))
	resultPane := widget.NewVScrollContainer(resultBox)
	content = *widget.NewVSplitContainer(statusPane, resultPane)
	window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			MenuItemFileClone(window.Canvas(), resultBox),
			MenuItemFilePull(window),
			MenuItemFileCheck(window),
			MenuItemFileOpen(window),
			fyne.NewMenuItemSeparator(),
			MenuItemFileExit(window),
		),
	))
	window.SetContent(&content)

	window.Resize(fyne.NewSize(1024, 768))
	window.ShowAndRun()
}
