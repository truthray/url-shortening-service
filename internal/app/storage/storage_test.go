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
	_, ok = storage.GetURL(0)
	assert.Equal(t, ok, false)

	storage.AddURL("testurl-1")
	assert.Equal(t, storage.CurrentIndex(), 0)
	currentElement, ok = storage.GetURL(storage.CurrentIndex())
	assert.Equal(t, ok, true)
	assert.Equal(t, currentElement, "testurl-1")

	storage.AddURL("testurl-2")
	assert.Equal(t, storage.CurrentIndex(), 1)
	currentElement, ok = storage.GetURL(storage.CurrentIndex())
	assert.Equal(t, ok, true)
	assert.Equal(t, currentElement, "testurl-2")

	currentElement, ok = storage.GetURL(storage.CurrentIndex() - 1)
	assert.Equal(t, ok, true)
	assert.Equal(t, currentElement, "testurl-1")
}
