
package main

import (
	"fmt"
	"reflect"
	"io"
	"os"
)

// whole collection 
type Vmap struct {
	cnt      int
    amaps map[string]*Amap
}

type Amap struct {
	parent  *Amap
    kids map[string]*Amap
	name  string
    vars map[string]*Var
}

// all the fun stuff a var can have 
// options  a slice of vars 
//         we need to keep then in order
// actions etc
type Extras struct {
	options *[]*Var
	actions *[]*Var
}


type SetFunc func (*Vmap, *Amap, *Var) int

func doSet(vm *Vmap, am *Amap, v *Var) int {
	ret := 0
	fmt.Println(" doSet running  [", am.name, v.name,"]")
    return ret
}

// basic var , can also have params and extras but we may keep actions and options as params.
type Var struct {
	valuedouble float64
	valuestring string
	valuevar *Var
	comp, name string
	vtype string
	params map[string]*Var	
	extras *Extras
	setFunc func (*Vmap, *Amap, *Var) int

}


func (vmap *Vmap)setAmap (comp, vname string) (bool, *Var, *Amap){
	//var res int
	//fmt.Println(comp)
	if vmap.amaps == nil { 
		vmap.amaps = make(map[string]*Amap)
	}
	//res = 0
	e, f := vmap.amaps[comp]
    if !f {
		e = new(Amap)
		e.name = comp
		vmap.amaps[comp] = e
		//vmap.amaps[comp].vars[vname] = v
	}
	v, f := e.vars[vname]
    if !f {
		e.vars = make(map[string]*Var)
		//v = new(Var)
		v = &Var{}
		e.vars[vname] = v
		v.comp = comp
		v.name = vname
		//v.vtype = reflect.TypeOf(val).Name()
		//fmt.Println(" new var type: addr ", reflect.TypeOf(val).Name(), v , comp, vname)

	}
 	return f, v, e
}

func (vmap *Vmap)setdVal (comp, vname string, val float64) *Var{
		//fmt.Println(" comp done type ", reflect.TypeOf(val))
	//fmt.Println(vname)
	f,v,e:= vmap.setAmap(comp, vname)
   	if !f {
		 v.vtype = reflect.TypeOf(val).Name()
		 v.setFunc = doSet
   	}
	v.valuedouble = float64(val)
	if v.setFunc != nil {
		v.setFunc(vmap,e,v)
	}
	return v
}

func (vmap* Vmap)getdVal(comp, vname string ) (float64, bool){
	if vmap.amaps == nil { 
		return float64(0), false
	}
	e, f := vmap.amaps[comp]
    if f {
		v, f := e.vars[vname]
		if f {
			return v.valuedouble, true 
		}
	}
	return float64(0), false
}


func (v *Var) getdVal() float64{
	return v.valuedouble
}

func (v *Var) setdVal( val float64) {
	//fmt.Println("setdval #1 ", v.comp , v.name, val)
	v.valuedouble = val
	v.vtype = reflect.TypeOf(val).Name()

	//fmt.Println("setdval #2", v, v.comp , v.name, val)
}

func (v *Var) setdParam( name string, val float64) {
	//fmt.Println("setdParam #1 ", name, val, v.params)
    if v.params == nil {
		//fmt.Println("setdParam #1  need param map")
		vp := make(map[string]*Var)
		v.params = vp
	}
	e, f := v.params[name]
    if !f {
		//fmt.Println("setdParam #1  need param name ", name)
		e = new(Var)
		v.params[name] = e
		//vmap.amaps[comp].vars[vname] = v
	}
	e.setdVal(val)
	e.vtype = reflect.TypeOf(val).Name()

}

func (v *Var) setcParam( name string, val string) {
	//fmt.Println("setdParam #1 ", name, val, v.params)
    if v.params == nil {
		//fmt.Println("setdParam #1  need param map")
		vp := make(map[string]*Var)
		v.params = vp
	}
	e, f := v.params[name]
    if !f {
		//fmt.Println("setdParam #1  need param name ", name)
		e = new(Var)
		e.name = name
		v.params[name] = e
		//vmap.amaps[comp].vars[vname] = v
	}
	e.setcVal(val)
	e.vtype = reflect.TypeOf(val).Name()

}

// no overloaded functions in go
func (v *Var) setcVal(val string) {
	//fmt.Println("setval string  #1 ", v.comp , v.name, val)
	v.valuestring = val
	v.vtype = reflect.TypeOf(val).Name()
}

func (v *Var) Show(mode string) string {
	var s string
	var tval string
	if v.vtype == "float64" {
		tval = fmt.Sprintf("%2.3f", v.valuedouble)
	} else if v.vtype == "string" {
		tval = fmt.Sprintf("\"%s\"", v.valuestring)
	} else {
		tval = fmt.Sprintf("vtype:[%s]", v.vtype)
	}
	if mode == "naked" {
		s = fmt.Sprintf("\"%s\":%s", v.name, tval)
	} else if mode == "clothed" {
		s = fmt.Sprintf("\"%s\":{\"value\":%s}", v.name, tval)
	} else if mode == "full" {

		s = fmt.Sprintf("\"%s\":{\"value\":%s", v.name, tval)
		if v.params != nil {
			for ni,vi := range v.params {
				vi.name = ni
				s+= ","
				s+=vi.Show("naked")
			}
		}
		if v.extras != nil {
			e:= v.extras
			o:= e.options
			if o != nil {
				op:= *e.options 
				len := len(op)
				var i int
				s+= fmt.Sprintf(",\"options\":[")

				for i=0; i<len; i++ {
					if i>0 { 
					 s += ",{"
					} else {
						s += "{"
					}
					s += op[i].Show("full")
					s += "}"

				}
				s+= fmt.Sprintf("]")

			}
		}
		
		s += fmt.Sprintf("}")
	} else {
		s = fmt.Sprintf("\"%s\":%2.3f", v.name, v.valuedouble)
	}

	return s
}

//
// no overloaded functions in go
func (v *Var) addOpt(val *Var) {
	//fmt.Println("addOpt #1 ", v.name, val)
	if v.extras == nil {
		e:= new(Extras)
		v.extras = e
	}
	e:=v.extras
	o:=e.options
	if o == nil {
		o:=make([]*Var, 1)
		o[0] = val
		e.options = &o
	}else{
		op:= *e.options 
		//fmt.Println("addOpt #2a  op", op, len(op))
		len := len(op)
		o2:=make([]*Var, (len+1))
		//	vopt.setcParam("val1","thisisval1")
        var i int

		for i=0; i<len; i++ {
			o2[i] = op[i]
		}
		o2[i] = val
		e.options = &o2
		//opx := op.append(op,val)
		//e.options = &opx

	}
	//op :=  *e.options
	//fmt.Println("addOpt #2 ", v, val)
	//fmt.Println("addOpt #2  op", op)
	//op = op.append(op,val)
}

// no overloaded functions in go
func (v *Var) setiVal(val int64) {
	//fmt.Println("setval int #1 ", v.comp , v.name, val)
	v.valuedouble = float64(val)
	v.vtype = reflect.TypeOf(val).Name()
	//fmt.Println("setval #2", v, v.comp , v.name, val)
}



func main() {
    //var dv float64 = float64(22) 
	//var dv2 float64 = float64(22) 
	var v1 *Var
	var vopt *Var
    var vop *Var
    vmap := new(Vmap)
    vmap.amaps = make(map[string]*Amap)
    e, f := vmap.amaps["/system/ess"]
    if !f {
        e = new(Amap)
    }
	v := Var{}
	v.name = "testFunc"
	v.valuedouble = 99
    e.vars = make(map[string]*Var)
    vmap.amaps["/system/ess"] = e
	vmap.amaps["/system/ess"].vars["CurrentVoltage"] = &v
	v.valuedouble = 98
	v.setFunc = doSet
    //vmap.amaps["aa"].vars["bb"].value = 99

	v.setFunc(vmap,e,&v)

    //fmt.Println(vmap)
    //fmt.Println(vmap.cnt, vmap.amaps["/system/ess"].vars["CurrentVoltage"].valuedouble)
	v1 = vmap.setdVal("/system/bms", "Status", 56)
	v1.setdParam("MaxStatus", 5.6)

	v1.setdParam("MaxStatus", 5.8)
	//dv,f = vmap.getdVal("/system/bms", "Status")

	//fmt.Println(" new value %s:%s,[{}]",
	    // v1,
		// v1.comp,
		// v1.name,
		// v1.valuedouble, " dv ", dv, f, v1)
	
	v1.setdVal(23.4)
	v1.setdVal(23.5)
	//dv2,f = vmap.getdVal("/system/bms", "Status")
	v1 = vmap.setdVal("/system/bms", "Status", 33)

	// fmt.Println(" another new value %s:%s,[{}]",
	// 	v1.getdVal(), 
	// 	v1,
	// 	v1.comp,
	// 	v1.name,
	// 	v1.valuedouble, " dv2 ", dv2, f)
	// 	//, vmap.amaps["/system/bms"].vars["Status"].value, "["
	//, vmap.amaps["/system/bms"].vars["Status"].value
	//, "]"
	//)
	vop  = new(Var)
	vop.setcParam("val1","thisisval1")
	vop.setdParam("val2",23.45)
	//fmt.Println(" vopt", vopt)
	//fmt.Println(" vop", vop)
	vopt = new(Var)
	vopt.name = "option0"
	vopt.setdVal(3.56)
	vop.addOpt(vopt)

	vopt = new(Var)
	vopt.name = "option1"
	vopt.setdVal(5.67)
	vop.addOpt(vopt)
	vopt = new(Var)
	vopt.setcVal("my number")

	vopt.name = "option2"

	vop.addOpt(vopt)
	vop.setdVal(12.345)
	vop.name = "myVop"

	s := vop.Show("clothed")
	
	io.WriteString(os.Stdout, s)
	io.WriteString(os.Stdout, "\n")

	s = vop.Show("full")
	
	io.WriteString(os.Stdout, s)
	io.WriteString(os.Stdout, "\n")

	// func := vop.addActFun("onSet","func")
	// func.addActItem(vop)

}
