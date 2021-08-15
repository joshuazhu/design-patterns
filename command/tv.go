package main

import "fmt"

type tv struct {
	isOn bool
}

func (t *tv) on() {
	fmt.Println("TV turned on")
	t.isOn = true
}

func (t *tv) off() {
	fmt.Println("TV turned off")
	t.isOn = false
}
