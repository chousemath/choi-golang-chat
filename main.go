package main

import (
	"fmt"
	"time"
)

func say(s string) {
	fmt.Print("\n", time.Now(), "\n >>> ", s, "\n\n")
}

func main() {
	say("Welcome to Choi-Golang-Chat~")
}
