package content

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
	svg "github.com/ajstarks/svgo"
	"github.com/jonesrussell/sober/components"
)

func Home() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		var buf bytes.Buffer
		canvas := svg.New(&buf)

		// Define the dimensions of the SVG canvas
		width := 800
		height := 650

		// Start the SVG canvas with the defined dimensions
		canvas.Startview(width, height, 0, 0, width, height)
		canvas.Title("Sobriety Roadmaps")
		canvas.Desc("A roadmap to sobriety")

		// Calculate the center position
		centerX := width / 2
		centerY := 0

		// Apply a translate transformation to center the content
		canvas.Gtransform(fmt.Sprintf("translate(%d, %d)", centerX, centerY))

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

		// End the group transformation
		canvas.Gend()

		canvas.End()
		_, err := w.Write(buf.Bytes())
		return err
	})
}
