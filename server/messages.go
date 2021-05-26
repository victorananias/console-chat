package main

import "time"

type Messages struct {
}

type Message struct {
	Sender      string
	Description string
	Date        string
}

var messagesList []Message = make([]Message, 0)

func (messages *Messages) Add(text, sender string) {
	message := Message{Description: text, Date: getCurrentTime()}
	if sender != "" {
		message.Sender = sender
	}
	messagesList = append(messagesList, message)
}

func (messages *Messages) GetAll() []Message {
	return messagesList
}

func getCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04")
}
