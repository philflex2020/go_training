package main

import (
	"fims"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

// var fimsReceive chan fims.FimsMsg
// var stateTicker, pubTicker *time.Ticker
var f fims.Fims
var x int

// func ReceiveChannelxx(c chan<- fims.FimsMsg, f fims.Fims) {
// 	//for f.connected {
// 	for {
// 		msg, err := f.Receive()
// 		if err != nil {
// 			log.Printf("Had an error while receiving on %s: %s\n", msg.Uri, err)
// 		}
// 		// fmt.Printf("about to put the msg on the channel %v\n", msg)
// 		c <- msg
// 	}
// }
func readBody() ([]byte, int, error) {
	// Looking for twins.json
	cpath := "test.json"
	count := 1
	var err error
	if len(os.Args) < 1 {
		fmt.Printf("Test file argument [1]  not found. Usage '%s <test_file> <test_count>'. Trying current working directory\n")
	}
	if len(os.Args) < 2 {
		fmt.Printf("Count argument [2]  not found. Usage '%s <test_file> <test_count>'. Setting count to 1\n")
	}
	if len(os.Args) > 1 {
		cpath = os.Args[1]
	}
	s := "Unknown"
	if len(os.Args) > 2 {
		s = os.Args[2]
		count, err = strconv.Atoi(s)
	}
	if err != nil {
		fmt.Printf("Couldn't decode count %s: %s \n", s, err)
		count = 1
	}
	_, xerr := os.Stat(cpath)
	//if os.IsNotExist(xerr) {
	fmt.Printf("Test file %s err : %s \n", cpath, xerr)
	err = xerr
	//}
	if err == nil {
		body, xerr := ioutil.ReadFile(cpath)
		if xerr != nil {
			fmt.Printf("Couldn't read the file %s: %s\n", cpath, xerr)
			err = xerr
		}
		//fmt.Printf(body)
		fmt.Println(body)
		fmt.Printf(" body type %T \n", body)

		return body, count, err
	}
	return nil, count, err
}

func init() {
	var err error
	f, err = fims.Connect("Go Program")
	if err != nil {
		log.Fatal("Unable to connect to FIMS server")
	}
	// fimsReceive = make(chan fims.FimsMsg)
	// f.Subscribe("/go_program")
	fmt.Println(" starting receive ")

	// fmt.Printf(" starting receive  fims %T \n", f)
	// go ReceiveChannelxx(fimsReceive, f)
	// stateTicker = time.NewTicker(100 * time.Millisecond)
	//pubTicker = time.NewTicker(500 * time.Millisecond)
}
func main() {
	// runch := make(chan int, 2)
	// donech := make(chan int)
	x := 0
	b := 0
	body, count, err := readBody()
	b = len(body)
	bs := string(body)
	fmt.Printf(" sending [%s] size [%d]\n", body, b)
	defer f.Close() // This makes sure the FIMS connection gets closed no matter how the program exits
	if err != nil {
		return
	}
	start := time.Now()
	for x < count {
		//fmt.Printf(" sending [%d] \n", x)
		f.Send(fims.FimsMsg{
			Method: "pub",
			Uri:    "/go_test",
			Body:   bs,
		})
		x += 1
	}
	duration := time.Since(start)
	dur := float32(float32(duration.Microseconds()) / 1000.0)
	total := (x * b)
	rate := float32(total) / (dur / 1000.0)
	fmt.Printf(" sent count %d size %d total_size %d time (mS) %f rate %f \n", x, b, total, dur, rate)
}

