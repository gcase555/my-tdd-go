package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	content := Content{"test": "this is just a test"}

	t.Run("lookup a word that exists", func(t *testing.T) {
		got, err := content.Search("test")
		want := "this is just a test"

		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("lookup a word that does not exist", func(t *testing.T) {
		got, err := content.Search("unknown")

		assert.Errorf(t, err, ErrNotFound.Error())
		assert.Empty(t, got)
	})
}

func TestAdd(t *testing.T) {
	content := Content{}
	content.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := content.Search("test")

	assert.NoError(t, err)

	assert.Equal(t, want, got)
}