package main

import (
	"fmt"
	"github.com/hen6003/flago"
	"strconv"
)

func main() {
	flags := []string{"r", "x", "u", "help"}

	flagsBool, _ := flago.CheckFlags(flags)

	if flagsBool[3] {
		fmt.Println("Convert decimal to hex.\n\nOptions:\n  -r:      reverse, hex to decimal\n  -x:      include '0x' before\n  -u:      uppercase the output\n  --help:  show this help")
	}

	var str string
	if flagsBool[1] {
		str = "%#"
	} else {
		str = "%"
	}

	if flagsBool[2] {
		str += "X"
	} else {
		str += "x"
	}

	if flagsBool[0] {
		for _, v := range flago.NonFlags() {
			vint, err := strconv.ParseInt(v, 16, 0)
			if err == nil {
				fmt.Printf("%d\n", vint)
			}
		}
	} else {
		for _, v := range flago.NonFlags() {
			vint, err := strconv.Atoi(v)
			if err == nil {
				fmt.Printf(str+"\n", vint)
			}
		}
	}
}
