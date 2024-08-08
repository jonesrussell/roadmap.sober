package handlers_test

import (
	"errors"
	"fullstackdev42/sober/handlers"
	"fullstackdev42/sober/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockPageService struct{}

func (m *MockPageService) GetWebpage(page string) (services.Webpage, error) {
	if page == "error" {
		return services.Webpage{}, errors.New("error getting webpage")
	}
	return services.Webpage{Title: "Test Title", Content: "Test Content"}, nil
}

func TestDefaultHandler(t *testing.T) {
	tests := []struct {
		name    string
		page    string
		wantErr bool
	}{
		{
			name:    "Home Success",
			page:    "home",
			wantErr: false,
		},
		{
			name:    "Community Success",
			page:    "community",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := &handlers.DefaultHandler{
				PageService: &MockPageService{},
			}

			var err error
			if tt.page == "home" {
				err = h.Home(c)
			} else if tt.page == "community" {
				err = h.Community(c)
			}

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), "Test Title")
			assert.Contains(t, rec.Body.String(), "Test Content")
		})
	}
}

func TestRenderPage(t *testing.T) {
	tests := []struct {
		name    string
		page    string
		wantErr bool
	}{
		{
			name:    "Home Page",
			page:    "home",
			wantErr: false,
		},
		{
			name:    "Community Page",
			page:    "community",
			wantErr: false,
		},
		{
			name:    "Error Case",
			page:    "error",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := &handlers.DefaultHandler{
				PageService: &MockPageService{},
			}

			err := h.RenderPage(c, tt.page)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Contains(t, rec.Body.String(), "Test Title")
				assert.Contains(t, rec.Body.String(), "Test Content")
			}
		})
	}
}
