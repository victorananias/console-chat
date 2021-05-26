package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const apiUrl string = "http://localhost:8080"

type Api struct {
}

func (api *Api) JoinAs(username string) error {
	json := []byte(`{ "username": "` + username + `" }`)
	_, err := http.Post(apiUrl+"/join", "application/json", bytes.NewBuffer(json))
	return err
}

func (api *Api) SendMessage(username, message string) error {
	json := []byte(`{ "username": "` + message + `", "message": "` + message + `" }`)
	_, err := http.Post(apiUrl+"/messages", "application/json", bytes.NewBuffer(json))
	return err
}

func (api *Api) GetMessages() ([]Message, error) {
	response, err := http.Get(apiUrl + "/messages")
	if err != nil {
		return nil, err
	}
	var messages []Message
	err = json.NewDecoder(response.Body).Decode(&messages)
	if err != nil {
		return nil, err
	}
	return messages, err
}
