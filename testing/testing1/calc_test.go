//+build unit

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1Calculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}
func Test2Calculate(t *testing.T) {
	assert.Equal(t, Calculate(2), -4)
	assert.NotNil(t, Calculate(3))
}
