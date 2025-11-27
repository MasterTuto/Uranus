package bootstrapper

import (
	"bytes"
	"fmt"
	"os"

	"github.com/MasterTuto/gaia/pkg/renderer"
	"github.com/MasterTuto/gaia/pkg/window"
	"golang.org/x/net/html"
)

type GaiaWriter struct {
}

func (gw GaiaWriter) Write(content []byte) (int, error) {
	window.Init()
	defer window.Quit()

	windowContext := window.Create(600, 600)
	defer windowContext.Destroy()

	renderContext, err := renderer.New(windowContext.Window)
	if err != nil {
		fmt.Println("Error on initiating the renderer")
		os.Exit(1)
	}

	rootNode = html.Parse(bytes.NewReader(content))

	elements := make([]renderer.LayoutedElement, 1)
	elements[0] = renderer.LayoutedElement{
		Type:   renderer.Rect,
		PosX:   0,
		PosY:   0,
		Height: 600,
		Width:  600,
		Element: renderer.Element{
			TextContent: "Eu amo meu amor",
			Style: renderer.ElementStyle{
				Color:           "#ff0000",
				BackgroundColor: "#f0f",
				FontFamily:      "font.ttf",
				FontSize:        32,
			},
		},
	}
	err = renderContext.Render(elements)
	if err != nil {
		fmt.Println("Erro ao renderizar")
		return 0, nil
	}

	window.ListenToEvents()
	return 0, nil
}
