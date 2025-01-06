package main

// this file has all code for running the UI
import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func Window() {
	// MyWindow = mw
	task := app.New()
	mv := task.NewWindow("Task 1 GUI")
	totalChars := widget.NewLabel("")
	avgSentLengthLabel := widget.NewLabel("")
	wordsCountLabel := widget.NewLabel("")
	mostOccurred := widget.NewLabel("")
	button := widget.NewButton("Upload File", func() {
		dialog.NewFileOpen(func(uri fyne.URIReadCloser, err error) {
			if err == nil && uri != nil {
				fileName := uri.URI().Path()
				content, _ := readFile(fileName)
				totalCharacters := fileCharacterCount(content)
				avgSent := averageSentenceLength(content)
				wordsCount := totalWordsCount(content)
				most := MostOccurredWord(content)
				totalChars.SetText(fmt.Sprintf("Toplam karakter sayısı : %v", totalCharacters))
				avgSentLengthLabel.SetText(fmt.Sprintf("Ortalama cümle uzunluğu : %v", avgSent))
				wordsCountLabel.SetText(fmt.Sprintf("Toplam kelime sayısı: %v", wordsCount))
				mostOccurred.SetText(fmt.Sprintf("En sık geçen kelim/sayı : %v", most))
			}
		}, mv).Show()
	})
	ResultContainer := container.NewVBox(
		button,
		totalChars,
		avgSentLengthLabel,
		wordsCountLabel,
		mostOccurred,
	)
	mv.Resize(fyne.NewSize(700, 700))
	mv.SetContent(ResultContainer)
	mv.ShowAndRun()
}
