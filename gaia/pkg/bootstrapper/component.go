package bootstrapper

import (
	"github.com/MasterTuto/gaia/pkg/renderer"
	"github.com/a-h/templ"
)

type Component interface {
	Get() templ.Component
	Style() renderer.ElementStyle
}
