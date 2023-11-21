package main

import (
	"fmt"
	"os"

	"github.com/MasterTuto/Uranus/pkg/renderer"
	"github.com/MasterTuto/Uranus/pkg/window"
)

func main() {
	window.Init()
	defer window.Quit()

	context := window.Create()
	defer context.Destroy()

	renderContext, err := renderer.New(context.Window)
	if err != nil {
		fmt.Println("Error on initiating the renderer")
		os.Exit(1)
	}

	elements := make([]renderer.LayoutedElement, 1)
	elements[0] = renderer.LayoutedElement{
		Type:   renderer.Text,
		PosX:   0,
		PosY:   0,
		Height: 100,
		Width:  100,
		Element: renderer.Element{
			TextContent: "Hello, World",
			Style: renderer.ElementStyle{
				Color:      "#fff",
				FontFamily: "font.ttf",
				FontSize:   32,
			},
		},
	}
	renderContext.Render(elements)

	window.ListenToEvents()
}
