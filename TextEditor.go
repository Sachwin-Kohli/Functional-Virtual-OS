package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var count int = 1

func showTextEditor() {
	a := app.New()

	w := a.NewWindow("Sachwin's Text Editor")

	w.Resize(fyne.NewSize(600, 400))

	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Sachwin's Text Editor"),
		),
	)

	content.Add(widget.NewButton("Add New File", func() {
		content.Add(widget.NewLabel("New File" + strconv.Itoa(count)))
		count++
	}))

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter Text")

	input.Resize(fyne.NewSize(450, 300))

	saveBtn := widget.NewButton("Save Text File", func() {
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)

				uc.Write(textData)

			}, w)

		saveFileDialog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")
		saveFileDialog.Show()
	})

	openBtn := widget.NewButton("Open Text File", func() {
		openFileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				ReadData, _ := ioutil.ReadAll(r)

				output := fyne.NewStaticResource("New File", ReadData)

				viewData := widget.NewMultiLineEntry()

				viewData.SetText(string(output.StaticContent))

				w := fyne.CurrentApp().NewWindow(
					string(output.StaticName))

				w.SetContent(container.NewScroll(viewData))

				w.Resize(fyne.NewSize(300, 300))
				w.Show()

			}, w)

		openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))

		openFileDialog.Show()
	})

	w.SetContent(
		container.NewVBox(
			content,
			input,

			container.NewHBox(
				saveBtn,
				openBtn,
			),
		),
	)

	w.ShowAndRun()
}
