package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	assert.Equal(t, want, got)
	assert.Equal(t, 4, spySleeper.Calls)
}
