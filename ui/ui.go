package ui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func mainWindow(app fyne.App) {
	window := app.NewWindow("Go Notes")
	header := widget.NewLabel("Go Notes Application v1.0")
	message := widget.NewLabel("")

	headerContainer := container.NewVBox(header)
	headerContainer.Layout = layout.NewCenterLayout()

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

	// notes container
	// notesCanvas := container.NewWithoutLayout(widget.NewLabel(""))

	// Sample notes
	notes := []string{
		"Note 1: Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		"Note 2: Consectetur adipiscing elit.",
		"Note 3: Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		// Add more notes as needed
	}

	// Create a container for holding the cards
	cardsContainer := container.NewVBox()

	// Add cards to the container
	for _, note := range notes {
		card := createNoteCard(note)
		cardsContainer.Add(card)
	}

	// Set the container alignment to the right
	cardsContainer.Layout = layout.NewVBoxLayout()
	// Create a scroll container to allow scrolling if there are too many notes
	scrollContainer := container.NewScroll(cardsContainer)
	scrollContainer.SetMinSize(fyne.NewSize(1280, 500))

	// bottom form container
	window.SetContent(container.NewVBox(toolbar, headerContainer, message, scrollContainer))

	window.Resize(fyne.NewSize(1280, 720))
	window.SetFixedSize(true)
	window.CenterOnScreen()
	window.Show()
}

func createNoteCard(note string) fyne.CanvasObject {
	card := widget.NewCard("", "", widget.NewLabel(note))
	card.Resize(fyne.NewSize(200, card.MinSize().Height))

	// Create a VScroll container for vertical scrolling
	scrollContainer := container.NewVScroll(card)

	// Set the horizontal scrolling policy to ScrollBarAlwaysOff
	scrollContainer.SetMinSize(fyne.NewSize(1280, 100))

	// Create a general container for the card
	cardContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), scrollContainer)

	return cardContainer
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
