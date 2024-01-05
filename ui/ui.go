package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)



func mainWindow(app fyne.App) {
	window := app.NewWindow("Go Notes")
	header := widget.NewLabel("Go Notes Application v1.0")
	message := widget.NewLabel("")

	headerContainer := container.NewVBox(header)
	headerContainer.Layout = layout.NewCenterLayout()

	apiContent := container.NewVBox(message)
	apiContent.Resize(fyne.NewSize(400, apiContent.MinSize().Height))
	apiContent.Layout = layout.NewCenterLayout()

	// toolbar

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	window.SetContent(container.NewVBox(toolbar, headerContainer, apiContent))

	window.Resize(fyne.NewSize(1280, 720))
	window.SetFixedSize(true)
	window.CenterOnScreen()
	window.Show()
}



func App() {
	app := app.NewWithID("go-notes")
	app.Settings().SetTheme(theme.LightTheme())

	icon, icon_error := fyne.LoadResourceFromPath("icons/app.jpg")
	if icon_error != nil {
		fmt.Println(icon_error.Error())
	}

	app.SetIcon(icon)
	mainWindow(app)
	app.Run()
}



