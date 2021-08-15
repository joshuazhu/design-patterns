package main

import "fmt"

type offCommand struct {
	device device
}

func (o *offCommand) execute() {
	fmt.Println("Execute off command")
	o.device.off()
}
