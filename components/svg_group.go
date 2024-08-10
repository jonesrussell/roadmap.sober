package components

import (
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

func measureTextWidth(text string, face font.Face) float64 {
	d := &font.Drawer{
		Face: face,
	}
	return float64(d.MeasureString(text)) / 64.0 // fixed.Int26_6 to float64
}

func SVGGroup(index int, title string) templ.Component {
	rectY := 33.62030742190215 + float64(index)*60
	textY := 58.920307421902145 + float64(index)*60

	// Load the font face
	face := basicfont.Face7x13

	// Measure the text width
	textWidth := measureTextWidth(title, face)
	padding := 20.0                    // Padding on each side
	rectWidth := textWidth + 2*padding // Add padding to the width

	// Calculate the center position
	centerX := -213.18656999416484 + rectWidth/2

	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := fmt.Fprintf(w, `
			<g data-type="topic" data-title="%s">
				<rect x="%f" y="%f" width="%f" height="46.3" rx="5" fill="#fdff00" stroke="black" stroke-width="2.7" style="--hover-color: #d6d700"></rect>
				<text x="%f" y="%f" text-anchor="middle" dominant-baseline="middle" font-size="17" fill="#000000">
					<tspan>%s</tspan>
				</text>
			</g>
		`, title, centerX-rectWidth/2, rectY, rectWidth, centerX, textY, title)
		return err
	})
}
