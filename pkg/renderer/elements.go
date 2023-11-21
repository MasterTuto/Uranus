package renderer

type Tag uint64

const (
	Div Tag = iota
	Span
	Img
)

type Element struct {
	Tag         Tag
	TextContent string
	Style       ElementStyle
}

type DivElement struct {
	Children Element
	Element
}
