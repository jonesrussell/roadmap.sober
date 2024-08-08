package services_test

import (
	"errors"
	"fullstackdev42/sober/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockFileReader struct{}

func (m *MockFileReader) ReadFile(name string) ([]byte, error) {
	if name == "./content/error.html" {
		return nil, errors.New("file not found")
	}
	return []byte("<h1>Test Content</h1>"), nil
}

func TestGetWebpage(t *testing.T) {
	tests := []struct {
		name     string
		pageName string
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "Success Case",
			pageName: "home",
			wantErr:  false,
		},
		{
			name:     "Error Case",
			pageName: "error",
			wantErr:  true,
			errMsg:   "file not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &services.PageServiceImpl{
				FileReader: &MockFileReader{},
			}

			page, err := ps.GetWebpage(tt.pageName)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.pageName, page.Title)
				assert.Equal(t, "<h1>Test Content</h1>", page.Content)
			}
		})
	}
}
