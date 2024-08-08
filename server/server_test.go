package server_test

import (
	"fullstackdev42/sober/server"
	"fullstackdev42/sober/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockPageService struct{}

func (m *MockPageService) GetWebpage(pageName string) (services.Webpage, error) {
	// For testing, you can return a dummy PageData and nil error.
	return services.Webpage{
		Title:   "Test Title",
		Content: "Test Content",
	}, nil
}

func TestNewServer(t *testing.T) {
	mockService := &MockPageService{}
	server := server.NewServer(mockService)

	tests := []struct {
		name     string
		method   string
		path     string
		expected int
	}{
		{
			name:     "Home Route",
			method:   http.MethodGet,
			path:     "/",
			expected: http.StatusOK,
		},
		{
			name:     "Community Route",
			method:   http.MethodGet,
			path:     "/community",
			expected: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rec := httptest.NewRecorder()

			// Serve the HTTP request
			server.Echo.ServeHTTP(rec, req)

			// Check the status code
			assert.Equal(t, tt.expected, rec.Code)
		})
	}
}
