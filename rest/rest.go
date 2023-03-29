package rest

import (
	"bytes"
	"net/http"
)

func Get(url string, bearerToken string) (*http.Response, string) {
	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + bearerToken

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, "Error creating request]n[ERROR] -" + err.Error()
	}
	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, "Error on invoking GET call client.\n[ERROR] -" + err.Error()
	}

	return resp, ""
}

func Post(url string, bearerToken string, jsonBytes []byte) (int, string) {

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + bearerToken

	// Create a new request using http
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return -1, "Error while creating the request for POST bytes:" + err.Error()
	}
	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil || resp.StatusCode > 299 {
		return resp.StatusCode, "Error on response.\n[ERROR] -" + resp.Status
	}

	//errStr := LogResponse(resp)

	return resp.StatusCode, ""
}
