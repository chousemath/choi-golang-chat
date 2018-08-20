package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type message struct {
	Name      string
	Content   string
	Timestamp string
}

func now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func say(msg message) {
	fmt.Printf("\n\t[%s]\n\t$%s: %s\n\n", msg.Timestamp, msg.Name, msg.Content)
}
func reverse(msgContent string) string {
	msgContent = strings.TrimSuffix(strings.Replace(msgContent, "$reverse", "", -1), " ")
	runes := []rune(msgContent)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name?: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	name = strings.TrimSuffix(name, "\n")
	msg := message{name, "I have entered the chatroom!", now()}
	say(msg)

	for true {
		msgContent, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		msg = message{name, strings.TrimSuffix(msgContent, "\n"), now()}
		if msg.Content == "exit" {
			break
		} else if strings.Contains(msg.Content, "$reverse") {
			msg.Content = reverse(msg.Content)
		}
		say(msg)
	}
}
