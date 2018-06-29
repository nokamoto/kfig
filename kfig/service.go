package main

import (
	"errors"
	"net/http"
	"fmt"
)

// Service represents an upstream service.
// https://docs.konghq.com/0.13.x/admin-api/#service-object
type Service struct {
	Name *string `yaml:"name" json:"name,omitempty"`

	Retries *int `yaml:"retries" json:"retries"`

	ConnectTimeout *int `yaml:"connect_timeout" json:"connect_timeout"`

	WriteTimeout *int `yaml:"write_timeout" json:"write_timeout"`

	ReadTimeout *int `yaml:"read_timeout" json:"read_timeout"`

	URL *string `yaml:"url" json:"url"`

	Present bool `json:"-"`
}

func (s Service) sprint() string {
	return prettyObj(s.Present, s)
}

// https://docs.konghq.com/0.13.x/admin-api/#add-service
// https://docs.konghq.com/0.13.x/admin-api/#update-service
func (s Service) createOrUpdate(api string) (string, error) {
	if s.Name == nil {
		return "", errors.New("name not found")
	}

	res, err := http.Get(fmt.Sprintf("%s/services/%s", api, *s.Name))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		name := *s.Name
		s.Name = nil
		return callPatch(fmt.Sprintf("%s/services/%s", api, name), []int{200}, s)
	}

	return callCreate(fmt.Sprintf("%s/services", api), []int{201, 409}, s)
}

// https://docs.konghq.com/0.13.x/admin-api/#delete-service
func (s Service) delete(api string) (string, error) {
	if s.Name == nil {
		return "", errors.New("name not found")
	}

	return callDelete(fmt.Sprintf("%s/services/%s", api, *s.Name), []int{204})
}