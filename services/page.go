package services

import (
	"fullstackdev42/sober/content"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Webpage struct {
	Title   string
	Content templ.Component
}

type PageService interface {
	GetWebpage(pageName string) (Webpage, error)
}

type PageServiceImpl struct {
	pages map[string]Webpage
}

var _ PageService = &PageServiceImpl{}

func NewPageService() *PageServiceImpl {
	return &PageServiceImpl{
		pages: map[string]Webpage{
			"home": {
				Title:   "Home",
				Content: content.Home(),
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
