package renderer

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Engine struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

func New(window *sdl.Window) (r *Engine, err error) {
	err = ttf.Init()

	if err != nil {
		return
	}

	sdlRenderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	var renderer = &Engine{
		window:   window,
		renderer: sdlRenderer,
	}
	return renderer, err
}

func (renderer *Engine) surface() *sdl.Surface {
	surface, err := renderer.window.GetSurface()
	if err != nil {
		panic(err)
	}
	return surface
}
