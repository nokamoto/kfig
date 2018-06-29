package main

import (
	"fmt"
)

// Consumer represents a consumer - or a user - of a Service.
// https://docs.konghq.com/0.13.x/admin-api/#consumer-object
type Consumer struct {
	Username *string `yaml:"username" json:"username"`

	CustomID *string `yaml:"custom_id" json:"custom_id"`

	Present bool `json:"-"`

	KeyAuths []KeyAuth `yaml:"key_auths" json:"-"`
}

func (c Consumer) sprint() string {
	s := ""
	if c.Username != nil {
		s += fmt.Sprintf(" username=%s", *c.Username)
	}
	if c.CustomID != nil {
		s += fmt.Sprintf(" custom_id=%s", *c.CustomID)
	}
	s += prettyBool(c.Present)
	return s
}

func (c Consumer) identifier() (string, error) {
	if c.Username != nil {
		return *c.Username, nil
	}
	if c.CustomID != nil {
		return *c.CustomID, nil
	}
	return "", fmt.Errorf("required: username or id")
}

// https://docs.konghq.com/0.13.x/admin-api/#create-consumer
func (c Consumer) create(api string) (string, error) {
	return callCreate(fmt.Sprintf("%s/consumers", api), []int{201, 409}, c)
}

// https://docs.konghq.com/0.13.x/admin-api/#delete-consumer
func (c Consumer) delete(api string) (string, error) {
	id, err := c.identifier()
	if err != nil {
		return "", err
	}
	return callDelete(fmt.Sprintf("%s/consumers/%s", api, id), []int{204})
}