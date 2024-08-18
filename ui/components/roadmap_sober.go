package components

import (
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
	svg "github.com/ajstarks/svgo"
)

// Constants for SVG dimensions and layout
const (
	width       = 360
	height      = 650
	groupHeight = 50
)

type Step struct {
	Text       string `json:"text"`
	Content    string `json:"content"`
	PanelTitle string `json:"panel_title"`
}

// RoadmapSober returns a templ.Component that renders a sobriety roadmap SVG
func RoadmapSober() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		canvas := svg.New(w)

		// Initialize SVG canvas
		canvas.Start(width, height, `style="width: 100%; height: 100%;"`)
		canvas.Title("Sobriety Roadmaps")
		canvas.Desc("A roadmap to sobriety")

		centerX := width / 2

		// Define the steps in the sobriety journey
		groups := []Step{
			{Text: "Admit the Problem", Content: "Content for Admit the Problem", PanelTitle: "Admit the Problem"},
			{Text: "Seek Help", Content: "Content for Seek Help", PanelTitle: "Seek Help"},
			{Text: "Detoxification", Content: "Content for Detoxification", PanelTitle: "Detoxification"},
			{Text: "Therapy and Counseling", Content: "Content for Therapy and Counseling", PanelTitle: "Therapy and Counseling"},
			{Text: "Support Groups", Content: "Content for Support Groups", PanelTitle: "Support Groups"},
			{Text: "Healthy Lifestyle", Content: "Content for Healthy Lifestyle", PanelTitle: "Healthy Lifestyle"},
			{Text: "Avoid Triggers", Content: "Content for Avoid Triggers", PanelTitle: "Avoid Triggers"},
			{Text: "Long-Term Maintenance", Content: "Content for Long-Term Maintenance", PanelTitle: "Long-Term Maintenance"},
			{Text: "Forgive Yourself", Content: "Content for Forgive Yourself", PanelTitle: "Forgive Yourself"},
			{Text: "Celebrate Milestones", Content: "Content for Celebrate Milestones", PanelTitle: "Celebrate Milestones"},
		}

		var prevMiddleBottomX, prevMiddleBottomY int
		for i, group := range groups {
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

// Button is assumed to be a function that returns the coordinates of a button
// This function needs to be implemented elsewhere or its signature might need to be adjusted
// func Button(canvas *svg.SVG, index int, text string, centerX int, id string) (int, int, int, int)
