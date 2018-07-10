package main

import (
	"errors"
	"fmt"
)

// Plugin represents a configuration that will be executed during the HTTP request/response lifecycle.
// https://docs.konghq.com/0.13.x/admin-api/#plugin-object
type Plugin struct {
	ID *string `yaml:"-" json:"id,omitempty"`

	ServiceID *string `yaml:"-" json:"service_id,omitempty"`

	ConsumerID *string `yaml:"-" json:"consumer_id,omitempty"`

	RouteID *string `yaml:"-" json:"route_id,omitempty"`

	Name string `yaml:"name" json:"name"`

	Config map[string]interface{} `yaml:"config" json:"config"`

	Enabled bool `yaml:"enabled" json:"enabled"`
}

func (p Plugin) sprint() string {
	return prettyJSONObj(p)
}

func (p Plugin) noIDs() Plugin {
	p.ID = nil
	p.ServiceID = nil
	p.ConsumerID = nil
	p.RouteID = nil
	return p
}

// (Enabling the plugin on a Service) https://docs.konghq.com/plugins/key-authentication
// https://docs.konghq.com/0.13.x/admin-api/#add-plugin
// https://docs.konghq.com/0.13.x/admin-api/#update-plugin
func (p Plugin) createOrUpdate(api string, service Service) (string, error) {
	if service.Name == nil {
		return "", errors.New("name not found")
	}

	plugins, err := service.plugins(api)
	if err != nil {
		return "", err
	}

	for _, plugin := range plugins {
		if plugin.Name == p.Name {
			if plugin.ID == nil {
				return "", errors.New("id not found")
			}
			
			p = p.noIDs()
			p.ServiceID = plugin.ServiceID

			if p.Config == nil {
				p.Config = make(map[string]interface{})
			}

			return callPatch(fmt.Sprintf("%s/plugins/%s", api, *plugin.ID), []int{200}, p)
		}
	}

	return callCreate(fmt.Sprintf("%s/services/%s/plugins", api, *service.Name), []int{201}, p.noIDs())
}