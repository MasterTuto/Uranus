package events

import "github.com/veandco/go-sdl2/sdl"

func HandleEvent(event sdl.Event) {
	event.GetType()
}
