package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jonesrussell/sober/services"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockPageService struct{}

func NewMockPageService() *MockPageService {
	return &MockPageService{}
}

func (m *MockPageService) GetWebpage(pageName string) (services.Webpage, error) {
	return services.Webpage{
		Title:   "Mock Page",
		Content: services.MockComponent("Mock Content"),
	}, nil
}

func (m *MockPageService) GetContentByID(id string) (string, error) {
	return "mock content", nil
}

func TestRenderPage(t *testing.T) {
	// Initialize a new Echo instance
	e := echo.New()

	// Create a new DefaultHandler
	h := &DefaultHandler{
		PageService: NewMockPageService(),
		BasePath:    "/test", // Set the base path here
	}

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Create a new Echo context
	c := e.NewContext(req, httptest.NewRecorder())

	// Call the RenderPage method
	err := h.RenderPage(c, services.MockComponent("Test Page"), "Test Title")

	// Assert that there was no error
	assert.NoError(t, err)
}
