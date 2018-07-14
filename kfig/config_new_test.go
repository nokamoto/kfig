package main

import (
	"testing"
)

func TestNewConfig_empty(t *testing.T) {
	c, err := NewConfig([]byte{})
	if err != nil {
		t.Fatal(err)
	}

	if len(c.Consumers) != 0 {
		t.Fatalf("len(%v) != 0", c.Consumers)
	}

	if len(c.Services) != 0 {
		t.Fatalf("len(%v) != 0", c.Services)
	}
}

func TestNewConfig_default(t *testing.T) {
	yaml := load("../test/default.yaml", t)

	c, err := NewConfig([]byte(yaml))
	if err != nil {
		t.Fatal(err)
	}

	if len(c.Consumers) != 2 {
		t.Fatalf("len(%v) != 2, %s", c.Consumers, yaml)
	}

	if len(c.Consumers[0].KeyAuths) != 2 {
		t.Fatalf("len(%v) != 2, %s", c.Consumers[0].KeyAuths, yaml)
	}

	if len(c.Consumers[1].KeyAuths) != 0 {
		t.Fatalf("len(%v) != 0, %s", c.Consumers[1].KeyAuths, yaml)
	}

	if len(c.Services) != 2 {
		t.Fatalf("len(%v) != 2, %s", c.Services, yaml)
	}

	if len(c.Services[0].Routes) != 1 {
		t.Fatalf("len(%v) != 1, %s", c.Services[0].Routes, yaml)
	}

	if len(c.Services[0].Plugins) != 1 {
		t.Fatalf("len(%v) != 1, %s", c.Services[0].Plugins, yaml)
	}

	if len(c.Services[1].Routes) != 0 {
		t.Fatalf("len(%v) != 0, %s", c.Services[1].Routes, yaml)
	}

	if len(c.Services[1].Plugins) != 0 {
		t.Fatalf("len(%v) != 0, %s", c.Services[1].Plugins, yaml)
	}
}