package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Sender  string
	Message string
	Date    string
}

var Messages []Message = make([]Message, 0)

func main() {
	registerServerRoutes()
	startServer()
}

func registerServerRoutes() {
	http.HandleFunc("/", func(responseWritter http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			storeMessage(responseWritter, request)
		case http.MethodGet:
			getMessages(responseWritter, request)
		default:
			responseWritter.Header().Set("Content-Type", "application/json")
			responseWritter.WriteHeader(http.StatusNotFound)
			responseWritter.Write([]byte(`{"message":"404 not found"}`))
		}
		logRequest(request.Method, "/")
	})
}

func startServer() {
	fmt.Println("Server listening at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func storeMessage(responseWritter http.ResponseWriter, request *http.Request) {
	responseWritter.Header().Set("Content-Type", "application/json")
	message := Message{}
	err := json.NewDecoder(request.Body).Decode(&message)
	if err != nil {
		http.Error(responseWritter, err.Error(), http.StatusBadRequest)
	}
	message.Date = time.Now().Format("yyyy-MM-dd HH:mm")
	Messages = append(Messages, message)
	responseWritter.WriteHeader(http.StatusOK)
	responseWritter.Write([]byte(`{"message": "Message sent."}`))
}

func getMessages(responseWritter http.ResponseWriter, request *http.Request) {
	responseWritter.Header().Set("Content-Type", "application/json")
	jsonResponse, _ := json.Marshal(Messages)
	responseWritter.WriteHeader(http.StatusOK)
	responseWritter.Write(jsonResponse)
}

func logRequest(method string, route string) {
	fmt.Println("\"" + method + "\" call to \"" + route + "\"")
}
