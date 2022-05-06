package main
// playing with things
import "fmt"
import "encoding/json"
//import "unsafe"

type Thing struct {
	name string
	id int
	ChargeRate float32
	DischargeRate float32
	MaxCapacity float32
	Capacity float32
	Output float32

}


func main() {

	var InputPower float32
	var PowerUsed float32
	InputPower = -200.0
 	
	ta := make([]*Thing,4)
	fmt.Printf("ta: ")
	fmt.Println(ta)
	fmt.Printf("ta [2]:")

	ta[0] = new(Thing)
	ta[0].name = "t0 name"
	ta[0].id = 1000
	ta[0].Capacity = 4.6
	ta[0].MaxCapacity = 100.0
	ta[0].ChargeRate = 5.0

	ta[1] = new(Thing)
	ta[1].name = "t1 name"
	ta[1].id = 1001
	ta[1].Capacity = 6.6
	ta[1].MaxCapacity = 200.0
	ta[1].ChargeRate = 10.0
	
	ta[2] = new(Thing)
	ta[2].name = "t2 name"
	ta[2].id = 1002
	ta[2].Capacity = 5.6
	ta[2].MaxCapacity = 50.0
	ta[2].ChargeRate = 10.0


	ta[3] = new(Thing)
	ta[3].name = "t3 name"
	ta[3].id = 1003
	ta[3].Capacity = 3.6
	ta[3].MaxCapacity = 150.0
	ta[3].ChargeRate = 7.0

	fmt.Println(ta[2])

	fmt.Printf("ta fixed up: \n")
	for _,tx := range ta {
		fmt.Println(tx)
	}

	// find max capacity
	maxCapacity:=ta[0].Capacity
	maxIx:=-1
	for ix,tx := range ta {
		if tx.Capacity > maxCapacity {
			maxCapacity = tx.Capacity
			maxIx = ix
		}
	}
	fmt.Printf("max capacity %v  : at %v \n", maxCapacity, maxIx)
	
	// find min capacity
	minCapacity:=ta[0].Capacity
	minIx:=-1
	for ix,tx := range ta {
		if tx.Capacity < minCapacity {
			minCapacity = tx.Capacity
			minIx = ix
		}
	}
	fmt.Printf("min capacity %v  : at %v \n", minCapacity, minIx)

	// Apply InputPower
	PowerUsed = InputPower
	for _,tx := range ta {
		if tx.Capacity < tx.MaxCapacity {
			if PowerUsed < 0 {
				// should be minimum of PowerUsed or chargerate
				fmt.Printf("tx  cycle: ")
				fmt.Println(tx)
				tx.Capacity += tx.ChargeRate
				PowerUsed += tx.ChargeRate
				fmt.Println(tx)
				if PowerUsed > 0 {
					tx.Capacity -= PowerUsed
					PowerUsed = 0
				}
			}
		}
	}
	fmt.Printf("PowerUsed %v  : out of  %v \n", PowerUsed, InputPower)
	fmt.Printf("ta after one cycle: ")
	fmt.Println(ta)
	fmt.Printf("ta after cycle : \n")
	for _,tx := range ta {
		fmt.Println(tx)
	}

	// bonus feature but we are headed in this direction
	tmap := make(map[string]interface{})
	tthings := make(map[string]interface{}) 
	tsystem := make(map[string]interface{})
	for _,tx := range ta {
		tmap[tx.name] = tx
	}	
	tsystem["/system"] = tthings
	tthings["things"] = tmap
	d,_ := json.Marshal(tsystem)
    fmt.Println(string(d))
}