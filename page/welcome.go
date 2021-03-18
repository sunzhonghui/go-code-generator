package page

import (
	"code/gen/resource/images"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func welcomeScreen(win fyne.Window) fyne.CanvasObject {
	logo := canvas.NewImageFromResource(images.ResourceIdmisstxPng)
	logo.FillMode = canvas.ImageFillContain
	logo.SetMinSize(fyne.NewSize(150, 150))

	return container.NewCenter(container.NewVBox(
		logo,
		container.NewHBox(
			widget.NewHyperlink("博客", parseURL("http://www.idmiss.com/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("文档", parseURL("http://www.idmiss.com/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("github", parseURL("http://www.idmiss.com/")),
		),
	))
}
