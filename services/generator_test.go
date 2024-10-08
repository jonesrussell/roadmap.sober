package services

import (
	"errors"
	"os"
	"testing"

	"github.com/jonesrussell/sober/content"
	"github.com/jonesrussell/sober/ui/components"
	"github.com/spf13/afero"
)

func (m *mockPageService) GetContentByID(id string) (string, error) {
	return "mock content", nil
}

func TestGeneratePage(t *testing.T) {
	// Create a new in-memory file system
	fs := afero.NewMemMapFs()

	// Create the StaticSiteService with the mock file system
	s := NewStaticSiteService(&mockPageService{})
	s.fs = fs

	step := &components.Step{
		Text:       "Admit the Problem",
		Content:    "Content for Admit the Problem",
		PanelTitle: "Admit the Problem",
	}

	err := s.generatePage("test.html", content.Home(step))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the file was created
	if _, err := fs.Stat("dist/test.html"); errors.Is(err, os.ErrNotExist) {
		t.Errorf("File was not created")
	}
}

func TestCopyPublicAssets(t *testing.T) {
	// Create a new in-memory file system
	fs := afero.NewMemMapFs()

	// Create the StaticSiteService with the mock file system
	s := &StaticSiteService{
		PageService: &mockPageService{},
		fs:          fs,
	}

	// Create the public directory and a test file in the mock file system
	fs.MkdirAll("public", 0755)
	afero.WriteFile(fs, "public/testfile.txt", []byte("test content"), 0644)

	// Run the copyPublicAssets method
	err := s.copyPublicAssets()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the file was created in the mock file system
	if _, err := fs.Stat("dist/static/testfile.txt"); errors.Is(err, os.ErrNotExist) {
		t.Errorf("File was not created")
	} else {
		content, _ := afero.ReadFile(fs, "dist/static/testfile.txt")
		if string(content) != "test content" {
			t.Errorf("File content was not copied correctly")
		}
	}
}
