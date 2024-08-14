package components

import (
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
	svg "github.com/ajstarks/svgo"
)

const (
	width       = 800
	height      = 650
	groupHeight = 50
)

func RoadmapSober() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		canvas := svg.New(w)

		// Start the SVG canvas with the defined dimensions
		canvas.Start(width, height, `style="width: 100%; height: 100%;"`)
		canvas.Title("Sobriety Roadmaps")
		canvas.Desc("A roadmap to sobriety")

		// Calculate the center position for horizontal centering
		centerX := width / 2

		groups := []string{
			"Admit the Problem",
			"Seek Help",
			"Detoxification",
			"Therapy and Counseling",
			"Support Groups",
			"Healthy Lifestyle",
			"Avoid Triggers",
			"Long-Term Maintenance",
			"Forgive Yourself",
			"Celebrate Milestones",
		}

		var prevMiddleBottomX, prevMiddleBottomY int
		for i := 0; i < len(groups); i++ {
			middleTopX, middleTopY, middleBottomX, middleBottomY := Button(canvas, i, groups[i], centerX, fmt.Sprintf("group-%d", i))

			if i > 0 {
				// Draw a vertical path from the previous group's bottom middle to the current group's top middle
				drawVerticalPath(canvas, prevMiddleBottomX, prevMiddleBottomY, middleTopX, middleTopY)
			}

			prevMiddleBottomX = middleBottomX
			prevMiddleBottomY = middleBottomY
		}

		canvas.End()

		return nil
	})
}

func drawVerticalPath(canvas *svg.SVG, startX, startY, endX, endY int) {
	path := fmt.Sprintf("M%d %d L%d %d", startX, startY, endX, endY)
	canvas.Path(path, "fill:none;stroke:black")
}
