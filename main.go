package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func say(name, message string) {
	fmt.Printf("\n[%s]\n$%s: %s\n\n", time.Now().Format("2006-01-02 15:04:05"), name, message)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name?: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	name = strings.TrimSuffix(name, "\n")
	say(name, "I have entered the chatroom!")
}
