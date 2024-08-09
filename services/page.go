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

type PageServiceImpl struct{}

var _ PageService = &PageServiceImpl{}

func (ps *PageServiceImpl) GetWebpage(pageName string) (Webpage, error) {
	var page Webpage
	switch pageName {
	case "home":
		page = Webpage{
			Title:   "Home",
			Content: content.Home(),
		}
	case "community":
		page = Webpage{
			Title:   "Community",
			Content: content.Community(),
		}
	default:
		return Webpage{}, echo.ErrNotFound
	}
	return page, nil
}
