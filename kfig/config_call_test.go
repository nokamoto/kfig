package main

import (
	"testing"
)

func TestCalls(t *testing.T) {
	yaml := load("../test/default.yaml", t)

	c, err := NewConfig([]byte(yaml))
	if err != nil {
		t.Fatal(err)
	}

	c.CallConsumers(*api)
	c.CallServices(*api)
}