package main
 
import (
	"fmt"
	. "strings"
)
type myReps struct {
	rep string
	with string
} 

func main() {
    jsonStr := `{
        "/system/cell_##CELL_NUM##" : {
			"Id","##CELL_ID##",
            "soc":50.2,
            "capacity", 5000
            "MaxChargeCurrent": -280,
            "MaxDischargeCurrent": 280,
            "state": "Normal"
        }
    }`
    newStr := jsonStr
	reps := []myReps{{"##CELL_ID##","MyCell_%02d"},{"##CELL_NUM##","%d"}}
	i :=2
	for _,x := range reps {
		fmt.Print(" >>Replace:", x.rep)
		y:= fmt.Sprintf(x.with,i)
    	fmt.Println(" >>With :",y)
		newStr = Replace(newStr,x.rep,y, -1)

	} 
	//newStr = Replace(newStr,"##CELL_NUM##","1", -1)
	// var x map[string]interface{}
 
    // json.Unmarshal([]byte(jsonStr), &x)
    // fmt.Println(x)

    // d,_ := json.Marshal(x)
    fmt.Println(newStr)


}
