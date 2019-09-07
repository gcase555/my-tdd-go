package structs_methods_interfaces

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	assert.Equal(t, want, got)
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 14.0}
		got := Area(rectangle)
		want := 168.00

		assert.Equal(t, want, got)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle(10)
		got := Area(circle)
		want := 314.1592653589793

		assert.Equal(t, want, got)
	})
}