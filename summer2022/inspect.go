import (
    "fmt"
    "reflect"
    "strings"
)

/*
InspectStruct prints the guts of an instantiated struct. Very handy for debugging
usage: InspectStruct(req, 0) -> prints all children
*/

func InspectStructV(val reflect.Value, level int) {
    if val.Kind() == reflect.Interface && !val.IsNil() {
        elm := val.Elem()
        if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
            val = elm
        }
    }
    if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }

    for i := 0; i < val.NumField(); i++ {
        valueField := val.Field(i)
        typeField := val.Type().Field(i)
        address := "not-addressable"

        if valueField.Kind() == reflect.Interface && !valueField.IsNil() {
            elm := valueField.Elem()
            if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
                valueField = elm
            }
        }

        if valueField.Kind() == reflect.Ptr {
            valueField = valueField.Elem()

        }
        if valueField.CanAddr() {
            address = fmt.Sprintf("0x%X", valueField.Addr().Pointer())
        }

        fmt.Printf("%vField Name: %s,\t Field Value: %v,\t Address: %v\t, Field type: %v\t, Field kind: %v\n",
            strings.Repeat("\t", level),
            typeField.Name,
            //valueField.Interface(),
            address,
            typeField.Type,
            valueField.Kind())

        if valueField.Kind() == reflect.Struct {
            InspectStructV(valueField, level+1)
        }
    }
}

func InspectStruct(v interface{}, level int) {
    InspectStructV(reflect.ValueOf(v), level)
}