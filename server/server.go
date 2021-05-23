package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type MessageRequest struct {
	Message  string
	Username string
}

type JoinChatRequest struct {
	Username string
}

type Message struct {
	Description string
	Date        string
}

var Messages []Message = make([]Message, 0)

func main() {
	registerServerRoutes()
	startServer()
}

func registerServerRoutes() {
	http.HandleFunc("/messages", func(responseWritter http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			storeMessage(responseWritter, request)
		case http.MethodGet:
			getMessages(responseWritter, request)
		default:
			notFoundResponse(responseWritter, request)
		}
		logRequest(request.Method, "/messages")
	})
	http.HandleFunc("/join", func(responseWritter http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			joinChat(responseWritter, request)
		default:
			notFoundResponse(responseWritter, request)
		}
		logRequest(request.Method, "/messages")
	})
}

func startServer() {
	fmt.Println("Server listening at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func storeMessage(responseWritter http.ResponseWriter, request *http.Request) {
	responseWritter.Header().Set("Content-Type", "application/json")
	requestBody := &MessageRequest{}
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(responseWritter, err.Error(), http.StatusBadRequest)
	}
	addMessage(requestBody.Username + ": " + requestBody.Message)
	responseWritter.WriteHeader(http.StatusOK)
	responseWritter.Write([]byte(`{"message": "Message sent."}`))
}

func getMessages(responseWritter http.ResponseWriter, request *http.Request) {
	responseWritter.Header().Set("Content-Type", "application/json")
	jsonResponse, _ := json.Marshal(Messages)
	responseWritter.WriteHeader(http.StatusOK)
	responseWritter.Write(jsonResponse)
}

func joinChat(responseWritter http.ResponseWriter, request *http.Request) {
	responseWritter.Header().Set("Content-Type", "application/json")
	requestBody := JoinChatRequest{}
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(responseWritter, err.Error(), http.StatusBadRequest)
	}
	addMessage(requestBody.Username + " joined the conversation.")
	responseWritter.WriteHeader(http.StatusOK)
	responseWritter.Write([]byte(`{"message": "Joined."}`))
}

func notFoundResponse(responseWritter http.ResponseWriter, request *http.Request) {
	responseWritter.Header().Set("Content-Type", "application/json")
	responseWritter.WriteHeader(http.StatusNotFound)
	responseWritter.Write([]byte(`{"message":"404 not found"}`))
}

func logRequest(method string, route string) {
	fmt.Println("\"" + method + "\" call to \"" + route + "\"")
}

func getCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04")
}

func addMessage(text string) {
	Messages = append(Messages, Message{Description: text, Date: getCurrentTime()})
}
