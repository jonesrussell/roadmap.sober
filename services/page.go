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

type PageServiceImpl struct {
	// Add your dependencies here...
}

var _ PageService = &PageServiceImpl{}

func (ps *PageServiceImpl) GetWebpage(pageName string) (Webpage, error) {
	// Read the content from an HTML file...
	content, err := os.ReadFile("./content/" + pageName + ".html")
	if err != nil {
		// Handle error...
		if os.IsNotExist(err) {
			return Webpage{}, err
		}
		return Webpage{}, err
	}

	return Webpage{
		Title:   pageName,
		Content: string(content),
	}, nil
}
