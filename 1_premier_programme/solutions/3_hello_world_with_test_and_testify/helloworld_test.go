package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_greet_should_return_the_name_prefixed_with_a_greeting(t *testing.T) {
	assert.Equal(t, "Hello, Jane", greet("Jane"))
}
