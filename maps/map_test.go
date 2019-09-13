package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	test_map := map[string]string{"test": "this is just a test"}

	got := Search(test_map, "test")
	want := "this is just a test"

	assert.Equal(t, want, got)
}
