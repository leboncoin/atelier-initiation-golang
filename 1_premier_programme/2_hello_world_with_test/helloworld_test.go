package main

import (
	"testing"
)

func Test_greet_should_return_the_name_prefixed_with_a_greeting(t *testing.T) {
	result := greet("Jane")
	if result != "Hello, Jane" {
		t.Errorf("expected '%v', got '%v'", "Hello, Jane", result)
	}
}
