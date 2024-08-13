package services

import (
	"context"
	"log"
	"os"

	"github.com/a-h/templ"
	"github.com/jonesrussell/sober/content"
	"github.com/jonesrussell/sober/ui/layout"
)

const (
	distDir      = "dist"
	publicDir    = "public/"
	staticOutput = "dist/static/"
)

type StaticSiteService struct {
	PageService PageService
}

func NewStaticSiteService(pageService PageService) *StaticSiteService {
	return &StaticSiteService{
		PageService: pageService,
	}
}

func (s *StaticSiteService) Generate(basePath string) error {
	// Delete dist directory if it exists
	if _, err := os.Stat("dist"); !os.IsNotExist(err) {
		os.RemoveAll("dist")
	}

	// Create dist directory
	os.Mkdir("dist", 0755)

	// Create and render home page
	homeContent := content.Home()
	if err := s.generatePage("index.html", layout.ContentPage(basePath, "Home", homeContent)); err != nil {
		log.Printf("Failed to generate home page: %v", err)
		return err
	}

	// Create and render community page
	communityContent := content.Community()
	if err := s.generatePage("community.html", layout.ContentPage(basePath, "Community", communityContent)); err != nil {
		log.Printf("Failed to generate community page: %v", err)
		return err
	}

	// Create and render 404 page
	notFoundContent := content.NotFound()
	if err := s.generatePage("404.html", layout.ContentPage(basePath, "404", notFoundContent)); err != nil {
		log.Printf("Failed to generate 404 page: %v", err)
		return err
	}

	// Copy public assets to dist/static directory
	if err := s.copyPublicAssets(); err != nil {
		log.Printf("Failed to copy public assets: %v", err)
		return err
	}

	return nil
}

func (s *StaticSiteService) generatePage(filename string, page templ.Component) error {
	file, err := os.Create(distDir + "/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := page.Render(context.Background(), file); err != nil {
		return err
	}

	return nil
}

func (s *StaticSiteService) copyPublicAssets() error {
	// Create dist/static directory
	if err := os.MkdirAll(staticOutput, 0755); err != nil {
		return err
	}

	inputFiles, err := os.ReadDir(publicDir)
	if err != nil {
		return err
	}

	for _, file := range inputFiles {
		inputFile := publicDir + file.Name()
		outputFile := staticOutput + file.Name()

		// Check if the file is a directory
		if file.IsDir() {
			// If it's a directory, create a corresponding directory in the output path
			if err := os.MkdirAll(outputFile, 0755); err != nil {
				return err
			}
		} else {
			// If it's a file, copy it
			input, err := os.ReadFile(inputFile)
			if err != nil {
				return err
			}

			if err := os.WriteFile(outputFile, input, 0644); err != nil {
				return err
			}
		}
	}

	return nil
}
