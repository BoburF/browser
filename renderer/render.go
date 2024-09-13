package renderer

import (
	"fmt"
	// "time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/widget"
)

func Render() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(700, 600))	

	file, err := fyne.LoadResourceFromPath("tel.jpg")
	if err != nil {
		fmt.Println("Image is not read!!!")
		return
	}
	image := canvas.NewImageFromResource(file)
	image.FillMode = canvas.ImageFillContain
	w.SetContent(image)

	// clock := widget.NewLabel("")
	// w.SetContent(widget.NewLabel("Hello World!"))
	// w.SetContent(clock)
	// go func() {
	// 	for range time.Tick(time.Second) {
	// 		updateTime(clock)
	// 	}
	// }()

	w.ShowAndRun()
	tidyUp()
}

// func updateTime(clock *widget.Label) {
// 	formatted := time.Now().Format("Time: 03:04:05")
// 	clock.SetText(formatted)
// }

func tidyUp() {
	fmt.Println("Exited")
}
