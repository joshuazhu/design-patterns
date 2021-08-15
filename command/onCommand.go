package main

import "fmt"

type onCommand struct {
	device device
}

func (o *onCommand) execute() {
	fmt.Println("Execute on command")
	o.device.on()
}
