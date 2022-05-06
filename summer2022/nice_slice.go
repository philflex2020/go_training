package main
// playing with slices
import "fmt"
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

    bs := "this is really a string"
    fmt.Println("bs:", bs)
	b := make([]byte,32)
	fmt.Printf("b [21]: %v\n", b[21])
	b[21] = 1
	fmt.Printf("b [21]: %v\n", b[21])

	fmt.Printf("&b [0]: %p\n", &b[0])
	fmt.Printf("&b [15]: %p\n", &b[15])

	b = []byte(bs)
	fmt.Println("byte:", b)

	
	ta := make([]Thing,4)
	fmt.Printf("ta: ")
	fmt.Println(ta)
	fmt.Printf("ta [2]:")


	ta[0].name = "t0 name"
	ta[0].id = 1000
	ta[0].Capacity = 4.6

	ta[1].name = "t1 name"
	ta[1].id = 1001
	ta[1].Capacity = 6.6
	
	ta[2].name = "t2 name"
	ta[2].id = 1002
	ta[2].Capacity = 5.6

	ta[3].name = "t3 name"
	ta[3].id = 1003
	ta[3].Capacity = 3.6
	fmt.Println(ta[2])

	fmt.Printf("ta fixed up: ")
	fmt.Println(ta)

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

	b2:=b[2:4]
	// fmt.Printf("b addr 1: %p\n", &b)
	// fmt.Printf("b addr 2: %p\n", (*byte)(unsafe.Pointer(&b)))
	// data:= &b[2]
	// fmt.Printf(" data: %v\n", data)
	// // data += 2
	// // xb := (*byte)(*data)
	// // fmt.Printf(" xb: %v\n", *xb)

	// fmt.Printf("b[2] addr: %p\n", (*byte)(unsafe.Pointer(&data)))
	fmt.Printf("b2: %v\n", b2)

	bs = string(b)
    fmt.Println("bs (string):", bs)
	bs = string(b2)
    fmt.Println("bs2 (string):", bs)

	s := make([]string, 3)
    fmt.Println("emp:", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}