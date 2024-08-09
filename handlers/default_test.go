package handlers

import (
	"fullstackdev42/sober/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRenderPage(t *testing.T) {
	// Initialize a new Echo instance
	e := echo.New()

	// Create a new DefaultHandler
	h := &DefaultHandler{
		PageService: services.NewMockPageService(),
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
