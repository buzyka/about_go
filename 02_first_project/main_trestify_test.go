package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddWithTestify(t *testing.T) {
	assert.Equal(t, 5, Add(2, 3), "2 + 3 should equal 5")
}
