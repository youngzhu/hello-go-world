package main

import (
	"flag"
	"fmt"
)

var stringVal string
var intVal int
var boolValue bool

func init() {
	flag.StringVar(&stringVal, "s", "", "string value")
	flag.BoolVar(&boolValue, "b", false, "bool value")
	flag.IntVar(&intVal, "i", -1, "int value")
}

// go run flag.go -s test -b true -i 10
// got: stringVal: test , boolValue: true , intVal: -1
// go run flag.go -s test -b false -i 10
// got: stringVal: test , boolValue: true , intVal: -1
// go run flag.go -b false -s test -i 10
// got: stringVal:  , boolValue: true , intVal: -1
// go run flag.go -s test -i 10 -b false
// got: stringVal: test , boolValue: true , intVal: 10
func main() {
	flag.Parse()

	fmt.Println("stringVal:", stringVal, ", boolValue:", boolValue, ", intVal:", intVal)
}