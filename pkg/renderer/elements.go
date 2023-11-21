package renderer

type Tag uint8
type LayoutedElementType uint8

const (
	Div Tag = iota
	Span
	Img
	TextNode
)

const (
	Rect LayoutedElementType = iota
	Text
	Image
)

type Element struct {
	Tag         Tag
	TextContent string
	Children    []*Element
	Attr        map[string]string
	Style       ElementStyle
}

type LayoutedElement struct {
	Type   LayoutedElementType
	PosX   int32
	PosY   int32
	Width  int32
	Height int32

	Element
}
