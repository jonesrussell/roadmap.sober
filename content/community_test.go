package content

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestCommunity(t *testing.T) {
	// Pipe the rendered template into goquery.
	r, w := io.Pipe()
	go func() {
		_ = Community().Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	// Expect the h1 element to be present and contain the correct text.
	h1 := doc.Find("h1").Text()
	expectedH1 := "Sober Friends"
	if h1 != expectedH1 {
		t.Errorf("expected h1 %q, got %q", expectedH1, h1)
	}

	// Expect the blockquote to be present and contain the correct text.
	blockquote := doc.Find("blockquote p").Text()
	expectedBlockquote := "If you hang out at a barbershop, sooner or later you're going to get a haircut."
	if blockquote != expectedBlockquote {
		t.Errorf("expected blockquote %q, got %q", expectedBlockquote, blockquote)
	}

	// Expect the citation to be present and contain the correct text.
	cite := doc.Find("cite").Text()
	expectedCite := "Somewhere"
	if cite != expectedCite {
		t.Errorf("expected cite %q, got %q", expectedCite, cite)
	}
}
