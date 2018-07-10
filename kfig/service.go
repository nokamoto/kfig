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

	Routes []Route `json:"-"`

	Plugins []Plugin `json:"-"`
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

// https://docs.konghq.com/0.13.x/admin-api/#list-routes-associated-to-a-service
func (s Service) routes(api string) ([]Route, error) {
	if s.Name == nil {
		return nil, errors.New("name not found")
	}

	var f func(string) ([]Route, error)

	f = func(url string) ([]Route, error) {
		retrive := retriveRoute{}

		if err := callGet(url, []int{200}, &retrive); err != nil {
			return nil, err
		}

		routes := retrive.Data

		if retrive.Next != nil {
			next, err := f(*retrive.Next)
			if err != nil {
				return nil, err
			}
			routes = append(routes, next...)
		}

		return routes, nil
	}

	return f(fmt.Sprintf("%s/services/%s/routes", api, *s.Name))
}

// https://docs.konghq.com/0.13.x/admin-api/#list-all-plugins
func (s Service) plugins(api string) ([]Plugin, error) {
	if s.Name == nil {
		return nil, errors.New("name not found")
	}

	retrive := retrivePlugin{}

	if err := callGet(fmt.Sprintf("%s/services/%s/plugins", api, *s.Name), []int{200}, &retrive); err != nil {
		return nil, err
	}

	return retrive.Data, nil
}