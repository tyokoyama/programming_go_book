// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//!+broadcaster
type client struct {
	Name string
	Ch chan string
}
//type client chan<- string // an outgoing message channel
var clients = make(map[client]bool) // all connected clients

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.Ch <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Ch)
			
			var str string
			for cli := range clients {
				str +=  cli.Name + ","
			}
			for cli := range clients {
				cli.Ch <- str
			}
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	var cli client
	cli.Ch = make(chan string)
	//ch := make(chan string) // outgoing client messages
	go clientWriter(conn, cli)

	who := conn.RemoteAddr().String()
	cli.Name = who
	cli.Ch <- "You are " + who
	messages <- who + " has arrived"
	
	var str string
	for cli := range clients {
		str +=  cli.Name + ","
	}
	messages <- str
	entering <- cli 

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, cli client) {
	for msg := range cli.Ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
