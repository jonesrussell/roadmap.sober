package services

import (
	"context"
	"log"
	"os"

	"github.com/a-h/templ"
	"github.com/jonesrussell/sober/content"
	"github.com/jonesrussell/sober/ui"
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
	homeContent := content.Home()
	s.generatePage("index.html", ui.ContentPage(basePath, "Home", homeContent))

	// Create and render community page
	communityContent := content.Community()
	s.generatePage("community.html", ui.ContentPage(basePath, "Community", communityContent))

	// Create and render 404 page
	notFoundContent := content.NotFound()
	s.generatePage("404.html", ui.ContentPage(basePath, "404", notFoundContent))

	// Copy public assets to dist/static directory
	s.copyPublicAssets()
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
