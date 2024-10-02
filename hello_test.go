package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, world!"
	actual := Hello()
	if actual != expected {
		t.Errorf("hello() = %v; want %v", actual, expected)
	}
}
