package main
 
import (
    "encoding/json"
    "fmt"
)
 
func main() {
    jsonStr := `{
        "fruits" : {
            "num":2,
            "tasty":true,
            "a": "apple",
            "b": "banana"
        },
        "colors" : {
            "r": "red",
            "g": "green"
        }
    }`
 
    var x map[string]interface{}
 
    json.Unmarshal([]byte(jsonStr), &x)
    fmt.Println(x)

    d,_ := json.Marshal(x)
    fmt.Println(string(d))


}
