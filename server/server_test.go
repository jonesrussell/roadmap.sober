package server

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/a-h/templ"
	"github.com/jonesrussell/loggo"
	"github.com/jonesrussell/sober/services"
	"github.com/stretchr/testify/assert"
)

// StringComponent returns a templ.Component that writes s to the provided io.Writer.
func StringComponent(s string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, s)
		return err
	})
}

type MockPageService struct{}

func (m *MockPageService) GetWebpage(title string) (services.Webpage, error) {
	return services.Webpage{
		Title:   "Mock Page Title",
		Content: StringComponent("Mock Page Content"),
	}, nil
}

func (m *MockPageService) GetContentByID(id string) (string, error) {
	return "mock content", nil
}

func TestNewServer(t *testing.T) {
	// Create a new server with the mock page service and the mock logger
	mockLogger := loggo.NewMockLogger()
	s := NewServer(&MockPageService{}, mockLogger)

	// Test the home route
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	s.Echo.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Mock Page Content")

	// Test the community route
	req = httptest.NewRequest(http.MethodGet, "/community", nil)
	rec = httptest.NewRecorder()
	s.Echo.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Mock Page Content")
}
