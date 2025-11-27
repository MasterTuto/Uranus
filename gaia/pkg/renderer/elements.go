package renderer

type Tag uint8

// Elements are just the elements to render
// There is actually just a parellell with
// original WebKit source code

const (
	Div Tag = iota
	Span
	Img
	TextNode
)

type Element struct {
	Tag         Tag
	TextContent string
	Children    []*Element
	Attr        ElementAttr
	Style       ElementStyle
}

// LayoutedElements are elements already processed by
// the layout phase, and therefore already has Size atributes

type LayoutedElementType uint8

const (
	Rect LayoutedElementType = iota
	Text
	Image
)

type LayoutedElement struct {
	Type   LayoutedElementType
	PosX   int32
	PosY   int32
	Width  int32
	Height int32

	Element
}
