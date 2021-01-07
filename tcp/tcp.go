package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "QUIT" || temp == "quit" {
			break
		}

		var ch chan string

		if temp == "RESET" || temp == "reset" {
			ch = reset()
		} else {
			ch = sumLen(temp)
		}

		c.Write([]byte(<-ch))
	}
	c.Close()
}

func startTCP() {
	l, err := net.Listen("tcp4", TCP_PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
