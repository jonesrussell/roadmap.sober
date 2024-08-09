package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWebpage(t *testing.T) {
	service := NewPageService()

	tests := []struct {
		name     string
		pageName string
		wantErr  bool
	}{
		{
			name:     "Test Home Page",
			pageName: "home",
			wantErr:  false,
		},
		{
			name:     "Test Community Page",
			pageName: "community",
			wantErr:  false,
		},
		{
			name:     "Test Non-Existent Page",
			pageName: "nonexistent",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.GetWebpage(tt.pageName)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
