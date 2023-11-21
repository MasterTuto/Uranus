package renderer

import (
	"github.com/MasterTuto/Uranus/pkg/colors"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func (e *Engine) renderText(element LayoutedElement) (err error) {
	fontFamily := element.Style.FontFamily

	var font *ttf.Font
	var color sdl.Color
	var text *sdl.Surface
	if font, err = ttf.OpenFont(fontFamily, element.Style.FontSize); err != nil {
		return
	}

	if color, err = colors.ParseToSdl(element.Style.Color); err != nil {
		return
	}

	if text, err = font.RenderUTF8Blended(element.TextContent, color); err != nil {
		return
	}
	defer text.Free()

	var position = &sdl.Rect{
		X: element.PosX,
		Y: element.PosY,
		H: element.Height,
		W: element.Width,
	}
	if err = text.Blit(nil, e.surface(), position); err != nil {
		return
	}

	return err
}
