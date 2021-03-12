package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const file string = "properties_test.ini"

func TestWrongConfigFile(t *testing.T) {
	reader := CreateReader("test.ini")
	_, err := reader.ReadGivenFileIntoMap()
	assert.NotNil(t, err, "wrong file, it does not exist")
}
func TestCorrectConfigFile(t *testing.T) {
	reader := CreateReader(file)
	_, err := reader.ReadGivenFileIntoMap()
	assert.Nil(t, err, "wrong file, it does not exist")
}
