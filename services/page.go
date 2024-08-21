package services

import (
	"github.com/jonesrussell/sober/content"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Webpage struct {
	Title   string
	Content templ.Component
}

type PageService interface {
	GetWebpage(pageName string) (Webpage, error)
	GetContentByID(id string) (string, error)
}

type PageServiceImpl struct {
	pages map[string]Webpage
}

var _ PageService = &PageServiceImpl{}

func NewPageService() *PageServiceImpl {
	homeStep := &Step{}

	return &PageServiceImpl{
		pages: map[string]Webpage{
			"home": {
				Title:   "Home",
				Content: content.Home(homeStep),
			},
			"community": {
				Title:   "Community",
				Content: content.Community(),
			},
			// Add more pages here as needed
		},
	}
}

func (ps *PageServiceImpl) GetWebpage(pageName string) (Webpage, error) {
	page, ok := ps.pages[pageName]
	if !ok {
		return Webpage{}, echo.ErrNotFound
	}
	return page, nil
}

func (ps *PageServiceImpl) GetContentByID(id string) (string, error) {
	return "foobar", nil
}
