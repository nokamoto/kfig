package main

import (
	"flag"
	"io/ioutil"
	"testing"
)

var api = flag.String("admin", "http://localhost:8001", "")
var mock = flag.String("mock", "http://localhost:8000", "")

func init() {
	flag.Parse()
}

func load(filename string, t *testing.T) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}