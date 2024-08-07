package services

type PageData struct {
	Title string
}

type PageService struct {
	// Add your dependencies here...
}

func (ps *PageService) GetPageData() PageData {
	// Add your logic here...
	return PageData{
		Title: "Home",
	}
}

// Add more methods for other services...
