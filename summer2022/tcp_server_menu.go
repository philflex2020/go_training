package main
// tcp_server.go

import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing

type Cmds struct {
	cmds map[string]Cmd
}

type Cmd struct {
	key, help string
	Func func (net.Conn, *Cmd, string, *Cmds) bool
}



func (c Cmds) addCmd (key,help string, f func (net.Conn, *Cmd, string, *Cmds) bool ) int {
	c.cmds[key]=Cmd{key:key, help:help, Func:f}
	return 0
}

func (c Cmds) runCmd (conn net.Conn, message string) bool  {
    b := false
	f := strings.Fields(message)
	fun,ok := c.cmds[f[0]]
	if ok {
	   b =  fun.Func(conn, &fun, message, &c)	
	} else {
		conn.Write([]byte( message +" not understood, try \"help\"  \n"))
	} 
	return b
}

func quitFunc(conn net.Conn, b *Cmd, c string, x *Cmds) bool {
	conn.Write([]byte( b.key +": closing connection\n"))
	return true
}

func helpFunc(conn net.Conn, b *Cmd, c string, x *Cmds) bool {
	for _,y := range x.cmds {
		conn.Write([]byte( y.key +":" + y.help + "\n"))
	}
	return false
}

func main() {

  fmt.Println("Launching server >>> telnet localhost 8081")
  cmds := new(Cmds)
  cmds.cmds = make(map[string]Cmd)
  cmds.addCmd("quit", " exit the service", quitFunc)
  cmds.addCmd("exit", " quit the service", quitFunc)
  cmds.addCmd("help", " show help",  helpFunc)

  // listen on all interfaces
  ln, _ := net.Listen("tcp", ":8081")

  // accept connection on port
  conn, _ := ln.Accept()
  // // this did 
  // run loop forever (or until ctrl-c)
  for done != true {
    // will listen for message to process ending in newline (\n)
    message, _ := bufio.NewReader(conn).ReadString('\n')
    // output message received
	//fmt.Print("Message Received:", string(message))
	if len(message) == 0  {
		done = true
	}
	done = cmds.runCmd(conn, message)

	// // this did not work
	// if message == "quit" {
	// 	conn.Write([]byte("bye from quit" + "\n"))
	// 	done = true
	// }

	// if f[0] == "quit" {
	// 	conn.Write([]byte("bye from f[0]" + "\n"))
	// 	done = true
	// }
	// if f[0] == "help" {
	// 	// for _,y := range cmds.cmds {
    // // this did 
	// 	// }
	// }
    // // this did 
	if message[0:4] == "quit" {
		conn.Write([]byte("bye again" + "\n"))
		done = true
	}
    
    // sample process for string received
    //newmessage := strings.ToUpper(message)
    // send new string back to client
    //conn.Write([]byte(newmessage + "\n"))
  }
}