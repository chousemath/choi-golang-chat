# Simple Chat Service

> This is a very simple chat service that is meant to be cross-platform, and used only from the terminal / command line. This is mainly just my effort to become more familiar with the Go programming language and its standard library.

# Running the application in development mode

```bash
$ go run main.go
```

# To do list

* Connect this application to Firebase for communication accross machines
* Add image/file upload/download capabilities
* Add simple, ASCII games within the application
* Add email support (sending the contents of chat session to someone via email)
* In-application weather and financial reports
* Support for inline mathematics
* Add unit tests

# Notes on Golang syntax

```bash
# Every Go program must start with a package declaration
# Executables are kinds of programs that can run directly from terminal
# Libraries are collections of code we package together to use in other programs
# import keyword is how we include code from other packages
# main is special because it's called when you execute the program
# Characters are represented by a byte (remember a byte is an integer)
```

# Notes on Golang tooling

```bash
# go run takes files, compiles into executable in temp dir, and runs program
$ go run main.go
```