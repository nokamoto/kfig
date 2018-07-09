package main

import (
	"fmt"
	"errors"
	"reflect"
	"encoding/json"
)

// Route defines rules to match client requests.
// https://docs.konghq.com/0.13.x/admin-api/#route-object
type Route struct {
	ID *string `yaml:"-" json:"id,omitempty"`

	Protocols []string `yaml:"protocols" json:"protocols"`

	Methods *[]string `yaml:"methods" json:"methods"`

	Hosts *[]string `yaml:"hosts" json:"hosts"`

	Paths *[]string `yaml:"paths" json:"paths"`

	StripPath *bool `yaml:"strip_path" json:"strip_path"`

	PreserveHost *bool `yaml:"preserve_host" json:"preserve_host"`
}

func (r Route) sprint() string {
	data, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (r Route) functionallyEqual(obj Route) bool {
	return reflect.DeepEqual(r.Protocols, obj.Protocols) && 
	reflect.DeepEqual(r.Methods, obj.Methods) && 
	reflect.DeepEqual(r.Hosts, obj.Hosts) && 
	reflect.DeepEqual(r.Paths, obj.Paths) && 
	reflect.DeepEqual(r.StripPath, obj.StripPath) && 
	reflect.DeepEqual(r.PreserveHost, obj.PreserveHost)
}

// https://docs.konghq.com/0.13.x/getting-started/configuring-a-service/#add-a-route-for-the-service
// https://docs.konghq.com/0.13.x/admin-api/#add-route
func (r Route) add(api string, service Service) (string, error) {
	if service.Name == nil {
		return "", errors.New("name not found")
	}
	return callCreate(fmt.Sprintf("%s/services/%s/routes", api, *service.Name), []int{201}, r)
}

// https://docs.konghq.com/0.13.x/admin-api/#add-route
func (r Route) remove(api string) (string, error) {
	if r.ID == nil {
		return "", errors.New("ID not found")
	}
	return callDelete(fmt.Sprintf("%s/routes/%s", api, *r.ID), []int{204})
}