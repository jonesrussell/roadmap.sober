package components

import (
	"bytes"
	"fmt"
	"io"

	svg "github.com/ajstarks/svgo"
)

func RoadmapSober(w io.Writer) error {
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

	SVGGroup(canvas, 0, "Admit the Problem")
	SVGGroup(canvas, 1, "Seek Help")
	SVGGroup(canvas, 2, "Detoxification")
	SVGGroup(canvas, 3, "Therapy and Counseling")
	SVGGroup(canvas, 4, "Support Groups")
	SVGGroup(canvas, 5, "Healthy Lifestyle")
	SVGGroup(canvas, 6, "Avoid Triggers")
	SVGGroup(canvas, 7, "Long-Term Maintenance")
	SVGGroup(canvas, 8, "Forgive Yourself")
	SVGGroup(canvas, 9, "Celebrate Milestones")

	// End the group transformation
	canvas.Gend()

	canvas.End()
	_, err := w.Write(buf.Bytes())
	return err
}
