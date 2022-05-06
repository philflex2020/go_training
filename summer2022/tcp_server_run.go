package main
// tcp_server.go

import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing
import "time"

type Cmds struct {
	cmds map[string]Cmd
	data map[string]string
}


type Cmd struct {
	key, help string
	Func func (net.Conn, *Cmd, *string, *Cmds) bool
}


func (c Cmds) addCmd (key,help string, f func (net.Conn, *Cmd, *string, *Cmds) bool ) int {
	c.cmds[key]=Cmd{key:key, help:help, Func:f}
	return 0
}

func (c Cmds) runCmd (conn net.Conn, message *string) bool  {
    b := false
	f := strings.Fields(*message)
	fun,ok := c.cmds[f[0]]
	if ok {
	   b =  fun.Func(conn, &fun, message, &c)	
	} else {
		conn.Write([]byte( *message +" not understood, try \"help\"  \n"))
	} 
	return b
}

func quitFunc(conn net.Conn, b *Cmd, c *string, x *Cmds) bool {
	conn.Write([]byte( b.key +": closing connection\n"))
	return true
}

func helpFunc(conn net.Conn, b *Cmd, c *string, x *Cmds) bool {
	for _,y := range x.cmds {
		conn.Write([]byte( y.key +":" + y.help + "\n"))
	}
	return false
}

func setFunc(conn net.Conn, b *Cmd, message *string, x *Cmds) bool {
	f := strings.Fields(*message)
	if len(f) > 2 {
		name := f[1]
		conn.Write([]byte( " name :" + name  ))
		value := f[2]
		conn.Write([]byte( " value :" + value  ))
		x.data[name]=value
	} else {
		conn.Write([]byte( "set needs two more items" ))
	}
	conn.Write([]byte( "\n" ))

	return false
}

func getFunc(conn net.Conn, b *Cmd, message *string, x *Cmds) bool {
	f := strings.Fields(*message)
	if len(f) > 1 {
		name := f[1]
		conn.Write([]byte( " name :" + name  ))

		value,ok := x.data[name]
		if ok {
			conn.Write([]byte( " value :" + value  ))
		} else {
			conn.Write([]byte( " value unknown " ))
		}
	} else {
		for x,y := range x.data {
			conn.Write([]byte( " name:" + x + " value:" + y + "\n"  ))
		}
	}
	conn.Write([]byte( "\n" ))

	return false
}

func run (cmds *Cmds, port string, done *bool) {
   p:= ":" + port

  fmt.Println("Launching server >>> telnet localhost" + p )
	// listen on all interfaces
  ln, _ := net.Listen("tcp", p)

  // accept connection on port
  conn, _ := ln.Accept()
  //fmt.Printf("Conn type %T:", conn)
  *done = false
  conn.Write([]byte("welcome from port " + port + "\n"))

  // run loop forever (or until ctrl-c)
  for *done != true {
    // will listen for message to process ending in newline (\n)
    message, _ := bufio.NewReader(conn).ReadString('\n')
    // output message received
	//fmt.Print("Message Received:", string(message))
	if len(message) == 0  {
		*done = true
	}
	*done = cmds.runCmd(conn, &message)

	
    // // this did 
	if message[0:4] == "quit" {
		conn.Write([]byte("bye again" + "\n"))
		*done = true
	}
  }
}

func main() {

  done := false
  cmds := new(Cmds)
  cmds.cmds = make(map[string]Cmd)
  cmds.data = make(map[string]string)
  cmds.addCmd("quit", " exit the service", quitFunc)
  cmds.addCmd("exit", " quit the service", quitFunc)
  cmds.addCmd("help", " show help",  helpFunc)
  cmds.addCmd("set",  " set name value - set a data value",  setFunc)
  cmds.addCmd("get",  " get name       - get a data value",  getFunc)
  go run(cmds, "8081", &done)
  go run(cmds, "8082", &done)
  for done != true {
	time.Sleep	( 1 * time.Second)  
  }

}