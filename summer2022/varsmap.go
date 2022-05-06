package main
//https://www.tutorialspoint.com/go/go_interfaces.htm
import (
	"fmt" 
	"math"
)

/* define an interface */
type AVal struct {
   valuefloat float64
   valueint int32
   valuebool int32
   valuestring string
}

/* define an interface */
type Shape interface {
	area() float64
 }
 
/* define a circle */
type Circle struct {
   x,y,radius float64
}

/* define a rectangle */
type Rectangle struct {
   width, height float64
}

/* define a method for circle (implementation of Shape.area())*/
func(circle Circle) area() float64 {
   return math.Pi * circle.radius * circle.radius
}

/* define a method for rectangle (implementation of Shape.area())*/
func(rect Rectangle) area() float64 {
   return rect.width * rect.height
}

/* define a method for shape */
func getArea(shape Shape) float64 {
   return shape.area()
}
  
func main() {
	dataMap := map[string]map[string]AVal 
//    vmap := make([]map[string]map[string]AVal, 0)
//    //amap := make([]map[string]interface{}, 0)
//    dataMap = make(map[string]AVal)

    var dataVal AVal
    dataVal.valuefloat = 3.7

//    //data["CurrValue"]="3.56"
//    dataMap["MaxValue"] = dataVal
	dataMap["/status/rack_01"]["MinValue"] = dataVal
	fmt.Println(23)
   
//    //amap["Voltage"] = data
//    vmap["/status/rack_01"] = dataMap

//    circle := Circle{x:0,y:0,radius:5}
//    rectangle := Rectangle {width:10, height:5}
   
//    fmt.Printf("Circle area: %f\n",getArea(circle))
//    fmt.Printf("Rectangle area: %f\n",getArea(rectangle))
}
