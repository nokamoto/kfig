package main

import (
	"os"
	"fmt"
)

// Config for Kong Admin API.
// https://docs.konghq.com/0.13.x/admin-api
type Config struct {
	Consumers []Consumer
	
	Services []Service
}

func handleCall(code string, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println(code)
	}
}

func (c Config) callConsumers(api string) {
	for i, consumer := range c.Consumers {
		fmt.Printf("[c%d]%s\n", i, consumer.sprint())

		if consumer.Present {
			handleCall(consumer.create(api))

			for j, key := range consumer.KeyAuths {
				fmt.Printf("[c%dk%d]%s\n", i, j, key.sprint())

				if key.Present {
					handleCall(key.create(api, consumer))
				} else {
					handleCall(key.delete(api, consumer))
				}
			}

		} else {
			handleCall(consumer.delete(api))
		}
	}
}

func (c Config) callServices(api string) {
	for i, service := range c.Services {
		fmt.Printf("[s%d]%s\n", i, service.sprint())

		if service.Present {
			handleCall(service.createOrUpdate(api))

			routes, err := service.routes(api)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			added := addedRoutes(routes, service.Routes)
			removed := removedRoutes(routes, service.Routes)

			fmt.Printf("[s%d] %d routes, %d added %d removed\n", i, len(routes), len(added), len(removed))

			for j, route := range added {
				fmt.Printf("[s%dra%d]%s\n", i, j, route.sprint())
				handleCall(route.add(api, service))
			}

			for j, route := range removed {
				fmt.Printf("[s%drr%d]%s\n", i, j, route.sprint())
				handleCall(route.remove(api))
			}

		} else {
			handleCall(service.delete(api))
		}
	}
}