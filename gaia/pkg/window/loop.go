package window

import (
	"github.com/MasterTuto/gaia/pkg/events"
	"github.com/veandco/go-sdl2/sdl"
)

func Quit() {
	sdl.Quit()
}

func (context *WindowContext) Destroy() {
	context.Window.Destroy()
}

func Init() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
}

func Create(width, height int32) *WindowContext {
	window, err := sdl.CreateWindow("Uranus", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	return &WindowContext{
		Window: window,
	}
}

func ListenToEvents() {
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			default:
				events.HandleEvent(event)
			}
		}
	}
}
