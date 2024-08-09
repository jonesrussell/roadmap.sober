package services

import (
	"context"
	"log"
	"os"

	"github.com/a-h/templ"
	"github.com/jonesrussell/sober/components"
	"github.com/jonesrussell/sober/content"
)

type StaticSiteService struct {
	PageService PageService
}

func NewStaticSiteService(pageService PageService) *StaticSiteService {
	return &StaticSiteService{
		PageService: pageService,
	}
}

func (s *StaticSiteService) Generate(basePath string) {
	// Delete dist directory if it exists
	if _, err := os.Stat("dist"); !os.IsNotExist(err) {
		os.RemoveAll("dist")
	}

	// Create dist directory
	os.Mkdir("dist", 0755)

	// Create and render home page
	s.generatePage("index.html", components.ContentPage(basePath, "Home", content.Home()))

	// Create and render community page
	s.generatePage("community.html", components.ContentPage(basePath, "Community", content.Community()))

	// Create and render 404 page
	s.generatePage("404.html", components.ContentPage(basePath, "404", content.NotFound()))

	// Copy public assets to dist/static directory
	s.copyPublicAssets()

	log.Println("Static site generated successfully!")
}

func (s *StaticSiteService) generatePage(filename string, page templ.Component) {
	file, err := os.Create("dist/" + filename)
	if err != nil {
		log.Fatalf("failed to create %s: %v", filename, err)
	}
	defer file.Close()

	err = page.Render(context.Background(), file)
	if err != nil {
		log.Fatalf("failed to render %s: %v", filename, err)
	}
}

func (s *StaticSiteService) copyPublicAssets() {
	inputPath := "public/"
	outputPath := "dist/static/"

	// Create dist/static directory
	os.MkdirAll(outputPath, 0755)

	inputFiles, err := os.ReadDir(inputPath)
	if err != nil {
		log.Fatalf("failed to read public directory: %v", err)
	}

	for _, file := range inputFiles {
		inputFile := inputPath + file.Name()
		outputFile := outputPath + file.Name()

		input, err := os.ReadFile(inputFile)
		if err != nil {
			log.Fatalf("failed to read file %s: %v", inputFile, err)
		}

		err = os.WriteFile(outputFile, input, 0644)
		if err != nil {
			log.Fatalf("failed to write file %s: %v", outputFile, err)
		}
	}
}
