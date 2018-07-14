package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

// Config for Kong Admin API.
// https://docs.konghq.com/0.13.x/admin-api
type Config struct {
	Consumers []Consumer
	
	Services []Service
}

func handleCall(code string, err error) error {
	if err != nil {		
		return err
	}
	fmt.Println(code)
	return nil
}

// NewConfig makes Config from the yaml string.
func NewConfig(data []byte) (*Config, error) {
	config := Config{}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// CallConsumers calls consumer apis.
func (c *Config) CallConsumers(api string) error {
	var err error
	for i, consumer := range c.Consumers {
		fmt.Printf("[c%d]%s\n", i, consumer.sprint())

		if consumer.Present {
			err = handleCall(consumer.create(api))
			if err != nil {
				return err
			}

			for j, key := range consumer.KeyAuths {
				fmt.Printf("[c%dk%d]%s\n", i, j, key.sprint())

				if key.Present {
					err = handleCall(key.create(api, consumer))
					if err != nil {
						return err
					}
				} else {
					err = handleCall(key.delete(api, consumer))
					if err != nil {
						return err
					}
				}
			}

		} else {
			handleCall(consumer.delete(api))
		}
	}
	return nil
}

func (Config) callPlugins(api string, service Service, i int) error {
	for j, plugin := range service.Plugins {
		fmt.Printf("[s%dp%d]%s\n", i, j, plugin.sprint())
		err := handleCall(plugin.createOrUpdate(api, service))
		if err != nil {
			return err
		}
	}
	return nil
}

func (Config) callRoutes(api string, service Service, i int) error {
	routes, err := service.routes(api)
	if err != nil {
		return err
	}

	added := addedRoutes(routes, service.Routes)
	removed := removedRoutes(routes, service.Routes)

	fmt.Printf("[s%d] %d routes, %d added %d removed\n", i, len(routes), len(added), len(removed))

	for j, route := range added {
		fmt.Printf("[s%dra%d]%s\n", i, j, route.sprint())
		err = handleCall(route.add(api, service))
		if err != nil {
			return err
		}
	}

	for j, route := range removed {
		fmt.Printf("[s%drr%d]%s\n", i, j, route.sprint())
		err = handleCall(route.remove(api))
		if err != nil {
			return err
		}
	}

	return nil
}

// CallServices calls service apis.
func (c *Config) CallServices(api string) error {
	var err error
	for i, service := range c.Services {
		fmt.Printf("[s%d]%s\n", i, service.sprint())

		if service.Present {
			err = handleCall(service.createOrUpdate(api))
			if err != nil {
				return err
			}

			err = c.callRoutes(api, service, i)
			if err != nil {
				return err
			}

			err = c.callPlugins(api, service, i)
			if err != nil {
				return err
			}
		} else {
			err = handleCall(service.delete(api))
			if err != nil {
				return err
			}
		}
	}
	return nil
}