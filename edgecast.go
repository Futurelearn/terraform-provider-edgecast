package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type edgecast struct {
	apiKey  string
	baseURL string
}

func (e edgecast) Request(method string, path string, payload interface{}) (*http.Response, error) {

	client := &http.Client{}
	url := e.baseURL + path

	json, _ := json.Marshal(payload)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(json))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "TOK:"+e.apiKey)
	req.Header.Add("Host", "api.edgecast.com")

	if method == "POST" {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	return resp, err
}
