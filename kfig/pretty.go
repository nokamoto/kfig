package main

import (
	"encoding/json"
	"fmt"
)

func prettyBool(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func prettyObj(b bool, obj interface{}) string {
	data, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf(" present=%s %s", prettyBool(b), data)
}