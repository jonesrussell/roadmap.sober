// services/mock.go
package services

import (
	"context"
	"io"

	"github.com/stretchr/testify/mock"
)

type MockPageService struct {
	mock.Mock
}

func (m *MockPageService) GetWebpage(pageName string) (Webpage, error) {
	args := m.Called(pageName)
	return args.Get(0).(Webpage), args.Error(1)
}

// handlers/default_test.go and handlers/pages_test.go
func NewMockPageService() *MockPageService {
	return &MockPageService{}
}

type MockComponent string

func (m MockComponent) Render(ctx context.Context, w io.Writer) error {
	_, err := io.WriteString(w, string(m))
	return err
}
