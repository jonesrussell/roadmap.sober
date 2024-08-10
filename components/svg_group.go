package components

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
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

func SVGGroup(index int, title string) templ.Component {
	rectY := 33.62030742190215 + float64(index)*60
	textY := 58.920307421902145 + float64(index)*60

	// Load the font face
	face := basicfont.Face7x13

	// Measure the text width
	textWidth := measureTextWidth(title, face)
	padding := 10.0                    // Padding on each side
	rectWidth := textWidth + 2*padding // Add padding to the width

	// Calculate the center position
	centerX := -213.18656999416484 + rectWidth/2

	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		var buf bytes.Buffer
		canvas := svg.New(&buf)

		canvas.Gtransform(fmt.Sprintf("translate(%f, %f)", centerX-rectWidth/2, rectY))
		canvas.Rect(0, 0, int(rectWidth), 46, "rx=5", "fill=#fdff00", "stroke=black", "stroke-width=2.7", "style=--hover-color: #d6d700")
		canvas.Text(int(centerX-(centerX-rectWidth/2)), int(textY-rectY), title, "text-anchor=middle", "dominant-baseline=middle", "font-size=17", "fill=#000000")
		canvas.Gend()

		_, err := w.Write(buf.Bytes())
		return err
	})
}
