package main

import (
	"fmt"
	"io/ioutil"
	"bytes"
	"net/http"
	"encoding/json"
)

type responseHandler func([]int) (string, error)

func containStatusCode(ok []int, code int) bool {
	for _, x := range ok {
		if code == x {
			return true
		}
	}
	return false
}

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

		if !containStatusCode(ok, res.StatusCode) {
			return "", fmt.Errorf("%d - %s", res.StatusCode, string(body))	
		}
	
		return fmt.Sprintf("%d - %s", res.StatusCode, string(body)), nil
	}
}

func callCreate(api string, ok []int, obj interface{}) (string, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return handleResponse(http.Post(api, "application/json", bytes.NewBuffer(data)))(ok)
}

func callPatch(api string, ok []int, obj interface{}) (string, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("PATCH", api, bytes.NewBuffer(data))
	if err != nil {
		return "", nil
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	return handleResponse(client.Do(req))(ok)
}

func callDelete(api string, ok []int) (string, error) {
	req, err := http.NewRequest("DELETE", api, nil)
	if err != nil {
		return "", nil
	}

	client := &http.Client{}
	return handleResponse(client.Do(req))(ok)
}

func callGet(api string, ok []int, data interface{}) error {
	res, err := http.Get(api)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if !containStatusCode(ok, res.StatusCode) {
		return fmt.Errorf("%d - %s", res.StatusCode, string(body))
	}

	if err := json.Unmarshal(body, data); err != nil {
		return err
	}

	return nil
}