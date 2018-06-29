package main

import (
	"os"
	"fmt"
)

// Config for Kong Admin API.
// https://docs.konghq.com/0.13.x/admin-api
type Config struct {
	Consumers []Consumer
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