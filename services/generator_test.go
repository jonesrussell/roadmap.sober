package services

import (
	"os"
	"testing"

	"github.com/jonesrussell/sober/content"
)

func TestGeneratePage(t *testing.T) {
	// Create dist directory before running the test
	os.MkdirAll("dist", 0755)

	s := NewStaticSiteService(&mockPageService{})
	err := s.generatePage("test.html", content.Home())
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the file was created
	if _, err := os.Stat("dist/test.html"); os.IsNotExist(err) {
		t.Errorf("File was not created")
	} else {
		// Clean up after the test
		os.Remove("dist/test.html")
	}
}

func TestCopyPublicAssets(t *testing.T) {
	s := NewStaticSiteService(&mockPageService{})
	err := s.copyPublicAssets()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Add more checks here to verify the files were copied correctly
	// For example, you could check if a specific file exists in the dist/static directory
}
