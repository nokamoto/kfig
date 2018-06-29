package main

import (
	"fmt"
	"io/ioutil"
	"bytes"
	"net/http"
	"encoding/json"
)

type responseHandler func([]int) (string, error)

func handleResponse(res *http.Response, err error) responseHandler {
	return func(ok []int) (string, error) {
		if err != nil {
			return "", err
		}
		defer res.Body.Close()
	
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
	
		for _, status := range ok {
			if res.StatusCode == status {
				return fmt.Sprintf("%d - %s", res.StatusCode, string(body)), nil
			}
		}
	
		return "", fmt.Errorf("%d - %s", res.StatusCode, string(body))	
	}
}

func callCreate(api string, ok []int, obj interface{}) (string, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return handleResponse(http.Post(api, "application/json", bytes.NewBuffer(data)))(ok)
}

func callDelete(api string, ok []int) (string, error) {
	req, err := http.NewRequest("DELETE", api, nil)
	if err != nil {
		return "", nil
	}

	client := &http.Client{}
	return handleResponse(client.Do(req))(ok)
}