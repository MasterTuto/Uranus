package main

import (
	"github.com/MasterTuto/Uranus/pkg/renderer"
	"github.com/MasterTuto/Uranus/pkg/window"
)

func main() {
	window.Init()
	defer window.Quit()

	context := window.Create()
	defer context.Destroy()

	renderContext := renderer.New(context.Window)

	elements := make([]renderer.LayoutedElement, 1)
	elements[0] = renderer.LayoutedElement{
		Type:   renderer.Rect,
		PosX:   0,
		PosY:   0,
		Height: 100,
		Width:  100,
		Element: renderer.Element{
			Style: renderer.ElementStyle{
				BackgroundColor: "#ff0000",
			},
		},
	}
	renderContext.Render(elements)

	window.ListenToEvents()
}
