package main
// tcp_server.go
// package main
// import (
//     "fmt"
//     "time"
// )
// func main() {
// Tickers use a similar mechanism to timers: a channel that is sent values. Here we’ll use the select builtin on the channel to await the values as they arrive every 500ms.

//     ticker := time.NewTicker(500 * time.Millisecond)
//     done := make(chan bool)
//     go func() {
//         for {
//             select {
//             case <-done:
//                 return
//             case t := <-ticker.C:
//                 fmt.Println("Tick at", t)
//             }
//         }
//     }()
// Tickers can be stopped like timers. Once a ticker is stopped it won’t receive any more values on its channel. We’ll stop ours after 1600ms.

//     time.Sleep(1600 * time.Millisecond)
//     ticker.Stop()
//     done <- true
//     fmt.Println("Ticker stopped")
// }
import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing
import "time"

type Cmds struct {
	cmds map[string]Cmd
	data map[string]string 
	
	done chan bool

}


type Cmd struct {
	key, help string
	Func func (net.Conn, *Cmd, *string, *Cmds) bool
}


type Msg struct {
	conn net.Conn
	msg *string
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

func run (cmds *Cmds, port string, ch chan *Msg, done *bool) {
   p:= ":" + port

  fmt.Println("Launching server >>> telnet localhost" + p )
	// listen on all interfaces
  ln, _ := net.Listen("tcp", p)

  // accept connection on port
  conn, _ := ln.Accept()
  //fmt.Printf("Conn type %T:", conn)
  *done = false
  conn.Write([]byte("welcome from port " + port + "\n"))
  m := new(Msg)
  m.conn = conn
  // run loop forever (or until ctrl-c)
  for *done != true {
    // will listen for message to process ending in newline (\n)
    message, _ := bufio.NewReader(conn).ReadString('\n')
	m.msg = &message
	ch <-m
	// output message received
	//fmt.Print("Message Received:", string(message))
	if len(message) == 0  {
		*done = true
	}
//	*done = cmds.runCmd(conn, &message)

	
    // // this did 
	if message[0:4] == "quit" {
		conn.Write([]byte("bye again" + "\n"))
		*done = true
	}
	if *done {
		cmds.done <- true
	}
  }
}

func main() {

  done := false
  cmds := new(Cmds)
  cmds.cmds = make(map[string]Cmd)
  cmds.data = make(map[string]string)
  cmds.done = make(chan bool)

  cmds.addCmd("quit", " exit the service", quitFunc)
  cmds.addCmd("exit", " quit the service", quitFunc)
  cmds.addCmd("help", " show help",  helpFunc)
  cmds.addCmd("set",  " set name value - set a data value",  setFunc)
  cmds.addCmd("get",  " get name       - get a data value",  getFunc)
  ch0 := make(chan *Msg)
  ch1 := make(chan *Msg)
  go run(cmds, "8081", ch0, &done)
  go run(cmds, "8082", ch1, &done)
  ticker := time.NewTicker(500 * time.Millisecond)
  bms_init(cmds)

  go func() {
        for {
          select {
		  case <-cmds.done:
			fmt.Println(" Done at ", ticker.C)
			done = true
			return
        case t := <-ticker.C:
			fmt.Println("Tick at", t)
			bms_run(cmds)

		case c0 := <-ch0:
			fmt.Println("c0 ", c0)
			fmt.Println("c0.msg ", *c0.msg)
			done = cmds.runCmd(c0.conn, c0.msg)
		case c1 := <-ch1:
			fmt.Println("c1 ", c1)
			fmt.Println("c1.msg ", *c1.msg)
			done = cmds.runCmd(c1.conn, c1.msg)
			}
          }
      }()
  //Tickers can be stopped like timers. Once a ticker is stopped it won’t receive any more values on its channel. We’ll stop ours after 1600ms.
  
  	//time.Sleep(10000 * time.Millisecond)
  	//ticker.Stop()
  	//     done <- true
    //fmt.Println("Ticker stopped")
    for done != true {
	   time.Sleep	( 1 * time.Second)  
    }
}