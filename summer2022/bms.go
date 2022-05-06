package main
//https://www.tutorialspoint.com/go/go_maps.htm

import "fmt"

var bms_count int 

func bms_init(x *Cmds) {
	  fmt.Println("bms init \n")
	  bms_count = 0
	  v := fmt.Sprintf("%d", bms_count)
	  x.data["bms_count"]= v
	  x.idata["bms_count"]= bms_count
}
   
func bms_run(x *Cmds) {
	fmt.Print("bms run icount ")
	fmt.Print(x.idata["bms_count"])
	fmt.Println("  \n")
	bms_count++
	v := fmt.Sprintf("%d", bms_count)
	x.data["bms_count"]= v
	x.idata["bms_count"]= bms_count
}
