package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

type Message struct {
	Sender      string
	Description string
	Date        string
}

type Chat struct {
	messages []Message
}

var username string
var api Api

var cleaners map[string]func() = map[string]func(){
	"linux": func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
	"windows": func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
}

func (chat *Chat) Join() {
	chat.clear()
	fmt.Println("Please enter your name:")
	fmt.Print(">")
	username, _ = chat.readTypedText()
	api.JoinAs(username)
}

func (chat *Chat) Start() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for {
			chat.update()
		}
	}()
	go func() {
		for {
			chat.waitForNewMessage()
		}
	}()
	wg.Wait()
}

func (chat *Chat) waitForNewMessage() {
	message, _ := chat.readTypedText()
	chat.sendMessage(message)
}

func (chat *Chat) sendMessage(message string) {
	api.SendMessage(username, message)
}

func (chat *Chat) update() {
	shouldUpdate := chat.loadMessages()
	if !shouldUpdate {
		return
	}
	chat.clear()
	chat.renderMessages()
	fmt.Print("\n>")
}

func (chat *Chat) loadMessages() (shouldUpdate bool) {
	messages, err := api.GetMessages()
	if err != nil {
		return false
	}
	if len(messages) == len(chat.messages) {
		return false
	}
	chat.messages = messages
	return true
}

func (chat *Chat) renderMessages() {
	for _, message := range chat.messages {
		if message.Sender == "" {
			fmt.Println(message.Date + " " + message.Description)
		} else {
			sender := message.Sender
			if message.Sender == username {
				sender = "You"
			}
			fmt.Println(message.Date + " " + sender + ": " + message.Description)
		}
	}
}

func (chat *Chat) clear() {
	cleaners[runtime.GOOS]()
}

func (chat *Chat) readTypedText() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	return chat.removeNewLineFromText(text), err
}

func (chat *Chat) removeNewLineFromText(text string) string {
	text = strings.Replace(text, "\n", "", -1)
	return text
}
