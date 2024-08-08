package services

import (
	"os"
)

type Webpage struct {
	Title   string
	Content string
}

type PageService interface {
	GetWebpage(pageName string) (Webpage, error)
}

type FileReader interface {
	ReadFile(filename string) ([]byte, error)
}

type OSFileReader struct{}

func (osfr *OSFileReader) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

type PageServiceImpl struct {
	FileReader FileReader
}

var _ PageService = &PageServiceImpl{}

func (ps *PageServiceImpl) GetWebpage(pageName string) (Webpage, error) {
	// Read the content from an HTML file...
	content, err := ps.FileReader.ReadFile("./content/" + pageName + ".html")
	if err != nil {
		return Webpage{}, err
	}

	return Webpage{
		Title:   pageName,
		Content: string(content),
	}, nil
}
