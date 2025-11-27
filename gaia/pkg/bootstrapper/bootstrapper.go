package bootstrapper

import (
	"context"
)

func Bootstrap(c Component) {
	c.Get().Render(context.Background(), GaiaWriter{})
}
