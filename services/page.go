package services

import (
	"io/ioutil"
	"os"
)

type PageData struct {
	Title   string
	Content string
}

type PageService struct {
	// Add your dependencies here...
}

func (ps *PageService) GetPageData(pageName string) (PageData, error) {
	// Read the content from an HTML file...
	content, err := ioutil.ReadFile("./content/" + pageName + ".html")
	if err != nil {
		// Handle error...
		if os.IsNotExist(err) {
			return PageData{}, err
		}
		return PageData{}, err
	}

	return PageData{
		Title:   pageName,
		Content: string(content),
	}, nil
}
