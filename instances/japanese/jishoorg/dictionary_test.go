package jishoorg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionary(t *testing.T) {
	assert := assert.New(t)
	res, err := NewDictionary().Search("コード")
	assert.NoError(err)
	assert.Equal("コード", res.Data[0].Japanese[0].Reading)
}
