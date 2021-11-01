package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()

var myWindow fyne.Window = myApp.NewWindow("OS-X")

var btn1 fyne.Widget
var btn2 fyne.Widget

var img fyne.CanvasObject
var DeskBtn fyne.Widget

var panelContent *fyne.Container

func main() {

	day := widget.NewButton("light", func() {
		// light theme
		myApp.Settings().SetTheme(theme.LightTheme())
	})
	night := widget.NewButton("dark", func() {
		// Dark theme
		myApp.Settings().SetTheme(theme.DarkTheme())
	})

	img = canvas.NewImageFromFile("C:/Users/91817/Desktop/go-workspace/OS/space.jpg")

	btn1 = widget.NewButtonWithIcon("Weather App", theme.InfoIcon(), func() {
		showWeatherApp(myWindow)
	})

	btn2 = widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(), func() {
		showCalc()
	})

	DeskBtn = widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		myWindow.SetContent(container.NewBorder(nil, panelContent, nil, nil, img))
	})

	panelContent = container.NewVBox(
		container.NewVBox(
			container.NewGridWithColumns(1,
				container.NewGridWithColumns(5,
					DeskBtn,
					btn1,
					btn2,

					day,
					night,
				),
			)))

	myWindow.Resize(fyne.NewSize(1920, 1080))
	myWindow.CenterOnScreen()

	myWindow.SetContent(
		container.NewBorder(nil, panelContent, nil, nil, img),
	)
	myWindow.ShowAndRun()
}
