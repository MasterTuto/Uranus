package main

import (
	"github.com/MasterTuto/Uranus/web/components"
	"github.com/MasterTuto/gaia/pkg/bootstrapper"
)

func main() {
	bootstrapper.Bootstrap(components.App())
}
