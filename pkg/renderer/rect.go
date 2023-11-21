package renderer

import (
	"github.com/MasterTuto/Uranus/pkg/colors"
	"github.com/veandco/go-sdl2/sdl"
)

func (renderer *Renderer) renderRect(element LayoutedElement) {
	surface := renderer.surface()

	surface.FillRect(nil, 0)

	rect := sdl.Rect{
		X: element.PosX,
		Y: element.PosY,
		W: element.Width,
		H: element.Height,
	}

	color, err := colors.Parse(element.Style.BackgroundColor)

	if err != nil {
		panic(err)
	}

	colour := sdl.Color{R: color.R, G: color.G, B: color.B, A: color.A}
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)
}
