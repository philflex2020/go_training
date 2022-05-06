package main
 
import (
    "encoding/json"
    "fmt"
)


func main() {
 
	jsonStr := `{
		"/system":{
			"things":{
				"t0 name":{
					"ChargeRate":5,
					"DischargeRate":0,
					"MaxCapacity":100,
					"Capacity":9.6,
					"Output":0},
				"t1 name":{
					"ChargeRate":10,
					"DischargeRate":0,
					"MaxCapacity":200,
					"Capacity":16.6,
					"Output":0},
				"t2 name":{
					"ChargeRate":10,
					"DischargeRate":0,
					"MaxCapacity":50,
					"Capacity":15.6,
					"Output":0},
				"t3 name":{
					"ChargeRate":7,
					"DischargeRate":0,
					"MaxCapacity":150,
					"Capacity":10.6,
					"Output":0}
				}
			}
		}
	`
	var x map[string]interface{}
 
    fmt.Println(" decode into maps")
    json.Unmarshal([]byte(jsonStr), &x)
	fmt.Println(x)
	
	fmt.Println(" pick one" )
	xx := x["/system"] 

	fmt.Println( xx )
	fmt.Printf( " xx %T \n", xx )
	var mykey map[string]interface{}
	for key,val := range x {
		fmt.Printf( "key %s value %T \n", key,val)
		    
			for key2,val2 := range val.(map[string]interface{}) {
				if key2 == "things" {
					mykey = val2.(map[string]interface{})
				}
				fmt.Printf( "   key2 %s value2 %T \n", key2,val2)
				    for key3,val3 := range val2.(map[string]interface{}) {
					   fmt.Printf( "       key3 %s value3 %v \n", key3,val3)
				    } 
				}
	}
	fmt.Printf( " [things] type %T \n", mykey )
	fmt.Printf( " [things] var %v \n", mykey )
	fmt.Printf( " [things][t0 name] var %v \n", mykey["t0 name"] )   
	myt0 := mykey["t0 name"].(map[string]interface{})
	fmt.Printf( " [things][t0 name][Capacity] var %v \n", myt0["Capacity"] )
	var mytn = make(map[string]interface{})
	for key4,val4 := range myt0 { 
		mytn[key4] = val4
	}    
	//copy( mytn , myt0)
	fmt.Printf( " mytn %v \n", mytn )
	// change capacity
	mytn["Capacity"] = 55.2
	// insert into things
	mykey["newone"] = mytn


	//fmt.Printf( " xx [things] %T \n", xx["things"].(map[string]interface{}) )
	// for xy := xx {
	// 	fmt.Println(" pick xy" )
	// 	fmt.Println( xy )
	// }
	//xy := xx["things"]
	//fmt.Println( xy )
 

    fmt.Println(" encode into json")
    d,_ := json.Marshal(x)
    fmt.Println(string(d))


}
