package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"

	"github.com/fyne-io/terminal"
	"github.com/fyne-io/terminal/cmd/fyneterm/data"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fyne Terminal")
	w.SetPadded(false)

	bg := canvas.NewRectangle(&color.RGBA{8, 8, 8, 255})
	img := canvas.NewImageFromResource(data.FyneScene)
	img.FillMode = canvas.ImageFillContain
	img.Translucency = 0.85

	t := terminal.NewTerminal()
	w.SetContent(fyne.NewContainerWithLayout(layout.NewMaxLayout(), bg, img, t.BuildUI()))
	w.Resize(fyne.NewSize(420, 260))

	go func() {
		err := t.Run(w.Canvas())
		if err != nil {
			fyne.LogError("Failure in terminal", err)
		}
		a.Quit()
	}()
	w.ShowAndRun()
}