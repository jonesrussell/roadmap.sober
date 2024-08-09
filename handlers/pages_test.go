package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jonesrussell/sober/services"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	// Initialize a new Echo instance
	e := echo.New()

	// Create a new DefaultHandler
	mockPageService := services.NewMockPageService()
	h := &DefaultHandler{
		PageService: mockPageService,
	}

	// Set up expectations on the mockPageService
	mockPageService.On("GetWebpage", "home").Return(services.Webpage{Content: services.MockComponent("Test Content")}, nil)

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/home", nil)

	// Create a new Echo context
	c := e.NewContext(req, httptest.NewRecorder())

	// Call the Home method
	err := h.Home(c)

	// Assert that there was no error
	assert.NoError(t, err)
}

func TestCommunity(t *testing.T) {
	// Initialize a new Echo instance
	e := echo.New()

	// Create a new DefaultHandler
	mockPageService := services.NewMockPageService()
	h := &DefaultHandler{
		PageService: mockPageService,
	}

	// Set up expectations on the mockPageService
	mockPageService.On("GetWebpage", "community").Return(services.Webpage{Content: services.MockComponent("Test Content")}, nil)

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/community", nil)

	// Create a new Echo context
	c := e.NewContext(req, httptest.NewRecorder())

	// Call the Community method
	err := h.Community(c)

	// Assert that there was no error
	assert.NoError(t, err)
}
