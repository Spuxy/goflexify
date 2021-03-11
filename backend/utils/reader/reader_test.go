package reader

import (
	"testing"
)

type MockReader struct {
	ownConfig Config
	ownError  error
	filename  string
}

func (m *MockReader) ReadGivenFileIntoMap() (Config, error) {
	return m.ownConfig, m.ownError
}

func TestHello(t *testing.T) {
	mock := MockReader{}
}
