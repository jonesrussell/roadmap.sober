package content

import (
	"bytes"
	"context"
	"io"

	"github.com/a-h/templ"
	svg "github.com/ajstarks/svgo"
	"github.com/jonesrussell/sober/components"
)

type homeComponent struct{}

func (h homeComponent) Render(ctx context.Context, w io.Writer) error {
	var buf bytes.Buffer
	canvas := svg.New(&buf)

	canvas.Start(1118, 1631)
	canvas.Title("Sobriety Roadmaps")
	canvas.Desc("A roadmap to sobriety")

	components.SVGGroup(0, "Admit the Problem")
	components.SVGGroup(1, "Seek Help")
	components.SVGGroup(2, "Detoxification")
	components.SVGGroup(3, "Therapy and Counseling")
	components.SVGGroup(4, "Support Groups")
	components.SVGGroup(5, "Healthy Lifestyle")
	components.SVGGroup(6, "Avoid Triggers")
	components.SVGGroup(7, "Long-Term Maintenance")
	components.SVGGroup(8, "Forgive Yourself")
	components.SVGGroup(9, "Celebrate Milestones")

	canvas.End()
	_, err := w.Write(buf.Bytes())
	return err
}

func Home() templ.Component {
	return homeComponent{}
}
