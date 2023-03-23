package freedictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionary(t *testing.T) {
	assert := assert.New(t)
	res, err := NewDictionary().Search("code")
	assert.NoError(err)
	assert.Equal("code", res[0].Word)
}
