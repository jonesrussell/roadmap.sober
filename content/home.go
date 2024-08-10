package content

import (
	"bytes"
	"context"
	"io"

	"github.com/a-h/templ"
	svg "github.com/ajstarks/svgo"
	"github.com/jonesrussell/sober/components"
)

func Home() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		var buf bytes.Buffer
		canvas := svg.New(&buf)

		canvas.Startview(800, 650, 0, 0, 800, 650)
		canvas.Title("Sobriety Roadmaps")
		canvas.Desc("A roadmap to sobriety")

		components.SVGGroup(canvas, 0, "Admit the Problem")
		components.SVGGroup(canvas, 1, "Seek Help")
		components.SVGGroup(canvas, 2, "Detoxification")
		components.SVGGroup(canvas, 3, "Therapy and Counseling")
		components.SVGGroup(canvas, 4, "Support Groups")
		components.SVGGroup(canvas, 5, "Healthy Lifestyle")
		components.SVGGroup(canvas, 6, "Avoid Triggers")
		components.SVGGroup(canvas, 7, "Long-Term Maintenance")
		components.SVGGroup(canvas, 8, "Forgive Yourself")
		components.SVGGroup(canvas, 9, "Celebrate Milestones")

		canvas.End()
		_, err := w.Write(buf.Bytes())
		return err
	})
}
