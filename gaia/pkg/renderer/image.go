package renderer

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func (e *Engine) renderImage(element LayoutedElement) (err error) {
	image, err := img.Load(element.Attr.Src)
	if err != nil {
		return
	}
	defer image.Free()

	texture, err := e.renderer.CreateTextureFromSurface(image)
	if err != nil {
		return
	}
	defer texture.Destroy()

	src := sdl.Rect{
		X: 0,
		Y: 0,
		H: element.Height,
		W: element.Width,
	}
	dst := sdl.Rect{
		X: element.PosX,
		Y: element.PosY,
		W: element.Width,
		H: element.Height,
	}

	e.renderer.Copy(texture, &src, &dst)
	return
}
