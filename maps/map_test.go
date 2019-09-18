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
	t.Run("new word", func(t *testing.T) {
		content := Content{}
		content.Add("test", "this is just a test")

		want := "this is just a test"
		got, err := content.Search("test")

		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})
	t.Run("existing word", func(t *testing.T) {
			word := "test"
			definition := "this is just a test"
			content := Content{word: definition}
			err := content.Add(word, "new test")

			assert.Errorf(t, err, ErrWordExists.Error())
			assert.Equal(t, definition, content[word])
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		content := Content{word: definition}
		newDefinition := "new definition"
		content.Update(word, newDefinition)

		assert.Equal(t, newDefinition, content[word])
	})

	t.Run("new word", func(t *testing.T) {
			word := "test"
			definition := "this is just a test"
			content := Content{}
			err := content.Update(word, definition)
			assert.Errorf(t, err, ErrWordNotExists.Error())
	})
}
