package components

import (
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
	svg "github.com/ajstarks/svgo"
	"github.com/jonesrussell/sober/services" // Import the services package
)

// Constants for SVG dimensions and layout
const (
	width       = 360
	height      = 650
	groupHeight = 50
)

// RoadmapSober returns a templ.Component that renders a sobriety roadmap SVG
func RoadmapSober(stepService services.StepService) templ.Component { // Add stepService as a parameter
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		canvas := svg.New(w)

		// Initialize SVG canvas
		canvas.Start(width, height, `style="width: 100%; height: 100%;"`)
		canvas.Title("Sobriety Roadmaps")
		canvas.Desc("A roadmap to sobriety")

		centerX := width / 2

		// Get the steps from the StepService
		buttons := stepService.GetSteps()

		var prevMiddleBottomX, prevMiddleBottomY int
		for i, group := range buttons {
			middleTopX, middleTopY, middleBottomX, middleBottomY := Button(canvas, i, group.Text, centerX, fmt.Sprintf("group-%d", i))

			if i > 0 {
				drawVerticalPath(canvas, prevMiddleBottomX, prevMiddleBottomY, middleTopX, middleTopY)
			}

			prevMiddleBottomX, prevMiddleBottomY = middleBottomX, middleBottomY
		}

		canvas.End()

		return nil
	})
}

// drawVerticalPath draws a line between two points in the SVG
func drawVerticalPath(canvas *svg.SVG, startX, startY, endX, endY int) {
	path := fmt.Sprintf("M%d %d L%d %d", startX, startY, endX, endY)
	canvas.Path(path, "fill:none;stroke:black")
}
