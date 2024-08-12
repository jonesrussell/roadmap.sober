package content

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/jonesrussell/sober/components"
)

func Home() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return components.RoadmapSober(w)
	})
}
