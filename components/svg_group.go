package components

import (
	"fmt"

	svg "github.com/ajstarks/svgo"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

func measureTextWidth(text string, face font.Face) float64 {
	d := &font.Drawer{
		Face: face,
	}
	return float64(d.MeasureString(text)) / 64.0 // fixed.Int26_6 to float64
}

func SVGGroup(canvas *svg.SVG, index int, title string) (middleTopX, middleTopY, middleBottomX, middleBottomY int) {
	rectY := 33.62030742190215 + float64(index)*60
	textY := 58.920307421902145 + float64(index)*60

	// Load the font face
	face := basicfont.Face7x13

	// Measure the text width
	textWidth := measureTextWidth(title, face)
	padding := 20.0                    // Padding on each side
	rectWidth := textWidth + 2*padding // Add padding to the width

	// Calculate the center position
	centerX := 0 + rectWidth/2

	// Adjust the group position based on the calculated centerX and rectY
	canvas.Group(fmt.Sprintf("style=\"transform:translate(%fpx, %fpx);\"", centerX-rectWidth/2, rectY))
	canvas.Rect(0, 0, int(rectWidth), 46, "rx=5", "fill=#fdff00", "stroke=black", "stroke-width=2.7", "style=--hover-color: #d6d700")
	canvas.Text(int(centerX), int(textY-rectY), title, "text-anchor=middle", "dominant-baseline=middle", "font-size=17", "fill=#000000")
	canvas.Gend()

	// Calculate the middle points for anchoring lines
	middleTopX = int(centerX)
	middleTopY = int(rectY)
	middleBottomX = int(centerX)
	middleBottomY = int(rectY + 46)

	return middleTopX, middleTopY, middleBottomX, middleBottomY
}
