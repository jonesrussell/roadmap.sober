package components

import (
	"fmt"

	svg "github.com/ajstarks/svgo"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

const (
	padding    = 10
	rectHeight = 40
)

func measureTextWidth(text string, face font.Face) float64 {
	d := &font.Drawer{
		Face: face,
	}
	return float64(d.MeasureString(text)) / 64.0 // fixed.Int26_6 to float64
}

func Button(canvas *svg.SVG, index int, title string, centerX int, id string) (middleTopX, middleTopY, middleBottomX, middleBottomY int) {
	rectY := 33.6 + float64(index)*60
	textY := 58.9 + float64(index)*60

	// Load the font face
	face := basicfont.Face7x13

	// Measure the text width
	textWidth := measureTextWidth(title, face)
	rectWidth := textWidth + 2*padding // Add padding to the width

	// Calculate the center position for the group
	groupCenterX := centerX - int(rectWidth/2)

	// Adjust the group position based on the calculated groupCenterX and rectY
	canvas.Group(fmt.Sprintf(`id="%s" style="transform:translate(%fpx, %fpx); cursor: pointer;"`, id, float64(groupCenterX), rectY))
	canvas.Rect(0, 0, int(rectWidth), rectHeight, "rx=5", "fill=#fdff00", "stroke=black", "stroke-width=2.7", "style=--hover-color: #d6d700")
	canvas.Text(int(rectWidth/2), int(textY-rectY), title, "text-anchor=middle", "dominant-baseline=middle", "font-size=17", "fill=#000000")
	canvas.Gend()

	// Calculate the middle points for anchoring lines
	middleTopX = centerX
	middleTopY = int(rectY)
	middleBottomX = centerX
	middleBottomY = int(rectY + rectHeight)

	return middleTopX, middleTopY, middleBottomX, middleBottomY
}
