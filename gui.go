package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func loadImage(imagepath string) image.Image {
	file, err := os.Open(imagepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}
	return img
}

func loadCardImages(dirpath string) (cardimages [53]image.Image) {
	// Open directory
	f, err := os.Open(dirpath)
	if err != nil {
		log.Fatal(err)
	}
	// Read files to slice
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		card := parseCard(strings.Split(name, ".")[0])
		cardimages[card] = loadImage(dirpath + "/" + name)
	}

	return
}

func raise() {
	fmt.Println("raise")
}

func main() {
	cardImages := loadCardImages("resources")

	gameApp := app.New()
	window := gameApp.NewWindow("Holdem-up!")

	opCard1 := canvas.NewImageFromImage(cardImages[32])
	opCard1.FillMode = canvas.ImageFillOriginal
	opCard2 := canvas.NewImageFromImage(cardImages[0])
	opCard2.FillMode = canvas.ImageFillOriginal

	opCardsContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), opCard1, opCard2)
	topBox := fyne.NewContainerWithLayout(layout.NewCenterLayout(), opCardsContainer)

	tableCards := make([]*canvas.Image, 5)
	for i := range tableCards {
		tableCards[i] = canvas.NewImageFromImage(cardImages[52])
		tableCards[i].FillMode = canvas.ImageFillOriginal
	}

	tableCardsContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), tableCards[0], tableCards[1], tableCards[2], tableCards[3], tableCards[4])
	midBox := fyne.NewContainerWithLayout(layout.NewCenterLayout(), tableCardsContainer)

	plCard1 := canvas.NewImageFromImage(cardImages[43])
	plCard1.FillMode = canvas.ImageFillOriginal
	plCard2 := canvas.NewImageFromImage(cardImages[11])
	plCard2.FillMode = canvas.ImageFillOriginal

	buttonRaise := widget.NewButton("Raise", raise)
	buttonCall := widget.NewButton("Call / Check", raise)
	buttonFold := widget.NewButton("Fold", raise)
	buttonContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), buttonRaise, buttonCall, buttonFold)

	plCardsContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), plCard1, plCard2, widget.NewLabel("        "), buttonContainer)
	botBox := fyne.NewContainerWithLayout(layout.NewCenterLayout(), plCardsContainer)

	content := fyne.NewContainerWithLayout(layout.NewBorderLayout(topBox, botBox, nil, nil), topBox, botBox, midBox)
	window.SetContent(content)
	window.Resize(fyne.NewSize(960, 540))
	window.ShowAndRun()
}
