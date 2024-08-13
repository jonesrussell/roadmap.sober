package services

import (
	"github.com/jonesrussell/sober/content"
)

type mockPageService struct{}

func (m *mockPageService) GetWebpage(pageName string) (Webpage, error) {
	// For testing, you can return a static webpage regardless of the pageName.
	// In a more complex test, you might want to return different pages based on the pageName.
	return Webpage{
		Title:   "Test Page",
		Content: content.Home(),
	}, nil
}
