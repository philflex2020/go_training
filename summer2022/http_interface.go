package main

import (
	"net/http"
	"html/template"
	"fmt"
	"net"
	"errors"
)

var fruits = map[string]interface{}{
	 "Apple":  "apple",
	 "Orange": "orange",
	 "Pear":   "pear",
	 "Grape":  "grape",
}
func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
func Handler(w http.ResponseWriter, r *http.Request) {
   t, err := template.ParseFiles("interface_view.html") 
   if err != nil {
	 panic(err)
   }
   viewData := struct {
	Fruits map[string]interface{}
	}{
		fruits,
	}

   
   t.Execute(w, viewData)
}

func main() {
	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(ip)
	fmt.Printf(" navigate to \"http://%v:8080/view\" \n",ip)
  	http.HandleFunc("/view", Handler)
  	http.ListenAndServe(":8080", nil)
}
 
