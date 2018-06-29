package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Consumer represents a consumer - or a user - of a Service.
// https://docs.konghq.com/0.13.x/admin-api/#consumer-object
type Consumer struct {
	Username *string `yaml:"username" json:"username"`

	CustomID *string `yaml:"custom_id" json:"custom_id"`

	Present bool `json:"-"`
}

func (c Consumer) sprint() string {
	s := ""
	if c.Username != nil {
		s += fmt.Sprintf(" username=%s", *c.Username)
	}
	if c.CustomID != nil {
		s += fmt.Sprintf(" custom_id=%s", *c.CustomID)
	}
	if c.Present {
		s += " present=yes"
	} else {
		s += " present=no"
	}
	return s
}

// https://docs.konghq.com/0.13.x/admin-api/#create-consumer
func (c Consumer) create(api string) (string, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	res, err := http.Post(fmt.Sprintf("%s/consumers", api), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode != 201 && res.StatusCode != 409 {
		return "", fmt.Errorf("%d - %s", res.StatusCode, string(body))
	}

	return fmt.Sprintf("%d - %s", res.StatusCode, string(body)), nil
}

// https://docs.konghq.com/0.13.x/admin-api/#delete-consumer
func (c Consumer) delete(api string) (string, error) {
	id := ""
	if c.Username != nil {
		id = *c.Username
	} else if c.CustomID != nil {
		id = *c.CustomID
	} else {
		return "", fmt.Errorf("required: username or id")
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/consumers/%s", api, id), nil)
	if err != nil {
		return "", nil
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode != 204 {
		return "", fmt.Errorf("%d - %s", res.StatusCode, string(body))
	}

	return fmt.Sprintf("%d - %s", res.StatusCode, string(body)), nil
}
