package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	storage := New()
	var currentElement string
	var ok bool

	assert.Equal(t, storage.CurrentIndex(), -1)
	currentElement, ok = storage.GetUrl(0)
	assert.Equal(t, ok, false)

	storage.AddUrl("testurl-1")
	assert.Equal(t, storage.CurrentIndex(), 0)
	currentElement, ok = storage.GetUrl(storage.CurrentIndex())
	assert.Equal(t, ok, true)
	assert.Equal(t, currentElement, "testurl-1")

	storage.AddUrl("testurl-2")
	assert.Equal(t, storage.CurrentIndex(), 1)
	currentElement, ok = storage.GetUrl(storage.CurrentIndex())
	assert.Equal(t, ok, true)
	assert.Equal(t, currentElement, "testurl-2")

	currentElement, ok = storage.GetUrl(storage.CurrentIndex() - 1)
	assert.Equal(t, ok, true)
	assert.Equal(t, currentElement, "testurl-1")
}
