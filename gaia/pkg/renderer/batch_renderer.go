package renderer

func (renderer *Engine) Render(elements []LayoutedElement) error {
	for _, element := range elements {
		if err := renderer.renderElement(element); err != nil {
			return err
		}
	}

	return nil
}

func (renderer *Engine) renderElement(element LayoutedElement) error {
	defer renderer.window.UpdateSurface()

	switch element.Type {
	case Rect:
		renderer.renderRect(element)
		return nil
	case Text:
		return renderer.renderText(element)
	case Image:
		return renderer.renderImage(element)
	}
	return nil
}
