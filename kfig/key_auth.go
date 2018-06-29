package main

import (
	"fmt"
)

// KeyAuth for Key Authentication, also sometimes referred to as an API key.
// https://docs.konghq.com/plugins/key-authentication/
type KeyAuth struct {
	Key string `yaml:"key" json:"key"`

	Present bool `yaml:"present" json:"-"`
}

func (k KeyAuth) sprint() string {
	return fmt.Sprintf(" key=%s, present=%s", k.Key, prettyBool(k.Present))
}

// https://docs.konghq.com/plugins/key-authentication/#create-a-key
func (k KeyAuth) create(api string, consumer Consumer) (string, error) {
	id, err := consumer.identifier()
	if err != nil {
		return "", err
	}
	return callCreate(fmt.Sprintf("%s/consumers/%s/key-auth", api, id), []int{201, 409}, k)
}

// https://docs.konghq.com/plugins/key-authentication/#delete-a-key
func (k KeyAuth) delete(api string, consumer Consumer) (string, error) {
	id, err := consumer.identifier()
	if err != nil {
		return "", err
	}
	return callDelete(fmt.Sprintf("%s/consumers/%s/key-auth/%s", api, id, k.Key), []int{204, 404})
}