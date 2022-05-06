package main
// https://zetcode.com/golang/socket/
import (
    "fmt"
    "log"
    "net"
)

func main() {

    con, err := net.Dial("tcp", "172.19.57.71:17")

    checkErr(err)

    defer con.Close()

    msg := ""

    _, err = con.Write([]byte(msg))

    checkErr(err)

    reply := make([]byte, 1024)

    _, err = con.Read(reply)

    checkErr(err)

    fmt.Println(string(reply))
}

func checkErr(err error) {

    if err != nil {

        log.Fatal(err)
    }
}
