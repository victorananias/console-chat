package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Message struct {
	Sender  string
	Message string
	Date    string
}

type Chat struct {
	chat []string
}

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
	username, _ := chat.readTypedText()
}

func (chat *Chat) Start() {
	for {
		chat.update()
		chat.waitForNewMessage()
	}
}

func (chat *Chat) waitForNewMessage() {
	fmt.Print("\n>")
	message, _ := chat.readTypedText()
	chat.sendMessage(message)
}

func (chat *Chat) sendMessage(message string) {

}

func (chat *Chat) update() {
	chat.clear()
	chat.renderMessages()
}

func (chat *Chat) renderMessages() {
	for _, element := range chat.chat {
		fmt.Println(element)
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
