package renderer

import (
	"github.com/veandco/go-sdl2/sdl"
)

func New(window *sdl.Window) *Renderer {
	return &Renderer{
		window: window,
	}
}

func (renderer *Renderer) surface() *sdl.Surface {
	surface, err := renderer.window.GetSurface()
	if err != nil {
		panic(err)
	}
	return surface
}

func (renderer *Renderer) Render(elements []LayoutedElement) {
	for _, element := range elements {
		renderer.renderElement(element)
	}
}

func (renderer *Renderer) renderElement(element LayoutedElement) {
	switch element.Type {
	case Rect:
		renderer.renderRect(element)
	case Text:
		renderer.renderText(element)
	case Image:
		renderer.renderImage(element)
	}

	renderer.window.UpdateSurface()
}
