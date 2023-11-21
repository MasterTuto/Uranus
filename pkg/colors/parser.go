package colors

import (
	"fmt"
	"image/color"
	"strings"
)

func Parse(theColor string) (c color.RGBA, e error) {
	if strings.HasPrefix(theColor, "#") {
		return parseHex(theColor)
	}

	c.A = 0xFF
	e = fmt.Errorf("Invalid color")
	return
}

func parseHex(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
		c.A = 255
	case 9:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
		c.A = 255
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")
	}
	return
}
