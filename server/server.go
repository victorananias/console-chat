package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
}

type MessageRequest struct {
	Message  string
	Username string
}

type JoinChatRequest struct {
	Username string
}

var messages Messages

func (server *Server) registerRoutes() {
	http.HandleFunc("/messages", func(responseWritter http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			server.storeMessage(responseWritter, request)
		case http.MethodGet:
			server.getMessages(responseWritter, request)
		default:
			server.notFoundResponse(responseWritter, request)
		}
		server.logRequest(request.Method, "/messages")
	})

	http.HandleFunc("/join", func(responseWritter http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			server.joinChat(responseWritter, request)
		default:
			server.notFoundResponse(responseWritter, request)
		}
		server.logRequest(request.Method, "/join")
	})
}

func (server *Server) start() {
	fmt.Println("Server listening at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func (server *Server) storeMessage(responseWritter http.ResponseWriter, request *http.Request) {
	requestBody := &MessageRequest{}
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(responseWritter, err.Error(), http.StatusBadRequest)
	}
	messages.Add(requestBody.Message, requestBody.Username)
	server.respond(responseWritter, `{"message": "Message sent."}`, http.StatusOK)
}

func (server *Server) getMessages(responseWritter http.ResponseWriter, request *http.Request) {
	jsonResponse, _ := json.Marshal(messages.GetAll())
	server.respond(responseWritter, string(jsonResponse), http.StatusOK)
}

func (server *Server) joinChat(responseWritter http.ResponseWriter, request *http.Request) {
	var requestBody JoinChatRequest
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(responseWritter, err.Error(), http.StatusBadRequest)
	}
	messages.Add(requestBody.Username+" joined the conversation.", "")
	server.respond(responseWritter, string(`{"message": "Joined."}`), http.StatusOK)
}

func (server *Server) notFoundResponse(responseWritter http.ResponseWriter, request *http.Request) {
	server.respond(responseWritter, `{"message":"404 not found"}`, http.StatusNotFound)
}

func (server *Server) respond(responseWritter http.ResponseWriter, json string, status int) {
	responseWritter.Header().Set("Content-Type", "application/json")
	responseWritter.WriteHeader(http.StatusOK)
	responseWritter.Write([]byte(json))
}

func (server *Server) logRequest(method string, route string) {
	fmt.Println("\"" + method + "\" call to \"" + route + "\"")
}
