package main
// https://golangr.com/ https://golangr.com/socket-server/
// socket client for golang
// Very basic socket server

import "net"
import "fmt"
import "bufio"

func main() {
  fmt.Println("Start server...")

  // listen on port 8000
  ln, _ := net.Listen("tcp", ":8000")

  // accept connection
  conn, _ := ln.Accept()

  // run loop forever (or until ctrl-c)
  for {
    // get message, output
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message Received:", string(message))
  }
}
