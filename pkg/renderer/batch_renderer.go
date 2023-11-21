package renderer

func (renderer *Engine) Render(elements []LayoutedElement) {
	for _, element := range elements {
		renderer.renderElement(element)
	}
}

func (renderer *Engine) renderElement(element LayoutedElement) {
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
