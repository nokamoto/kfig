package main

import (
	"time"
	"io/ioutil"
	"bytes"
	"net/http"
	"testing"
)

func expect(res *http.Response, err error) func(*http.Request, int, *testing.T) {
	return func(req *http.Request, status int, t *testing.T) {
		if err != nil {
			t.Fatal(req, err)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(req, err)
		}

		if res.StatusCode != status {
			t.Fatalf("%v %d %s - %s", req, res.StatusCode, res.Status, string(body))
		}
	}
}

func TestCalls(t *testing.T) {
	yaml := load("../test/default.yaml", t)

	c, err := NewConfig([]byte(yaml))
	if err != nil {
		t.Fatal(err)
	}

	err = c.CallConsumers(*api)
	if err != nil {
		t.Fatal(err)
	}

	err = c.CallServices(*api)
	if err != nil {
		t.Fatal(err)
	}

	// wait for db_update_frequency
	// https://github.com/Kong/kong/blob/0.13.1/kong.conf.default#L388
	time.Sleep(5 * time.Second)

	client := &http.Client{}

	req, err := http.NewRequest("GET", *mock, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}
	req.Host = "example.com"

	expect(client.Do(req))(req, 401, t)

	req.Header.Add("apikey", "k0")
	expect(client.Do(req))(req, 200, t)

	req.Header.Set("apikey", "k1")
	expect(client.Do(req))(req, 403, t)
}