package main
// tcp_server.go

import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing


func main() {

  fmt.Println("Launching server >>> telnet localhost 8081")
  // listen on all interfaces
  ln, _ := net.Listen("tcp", ":8081")

  // accept connection on port
  conn, _ := ln.Accept()
  //fmt.Printf("Conn type %T:", conn)
  done := false
  // run loop forever (or until ctrl-c)
  for done != true {
    // will listen for message to process ending in newline (\n)
    message, _ := bufio.NewReader(conn).ReadString('\n')
    // output message received
	fmt.Print("Message Received:", string(message))
	if len(message) == 0  {
		done = true
	}
	// // this did not work
	// if message == "quit" {
	// 	conn.Write([]byte("bye from quit" + "\n"))
	// 	done = true
	// }

    // // this did 
	if message[0:4] == "quit" {
		conn.Write([]byte("bye again" + "\n"))
		done = true
	}
    
    // sample process for string received
    newmessage := strings.ToUpper(message)
    // send new string back to client
    conn.Write([]byte(newmessage + "\n"))
  }
}