package main

import (
	"fmt"
	"time"
)

type client struct {
	command command
}

func (c *client) invokeCommand() {
	c.command.execute()
}

type command interface {
	execute()
}

type orderCommand struct {
	chef chef
}

func (c *orderCommand) execute() {
	c.chef.cook()
}

type chef interface {
	cook()
}

type dessertChef struct {
	isCooking bool
}

func (c *dessertChef) cook() {
	c.isCooking = true
	fmt.Println("chef started cooking your dessert")
	time.Sleep(time.Second * 2)
	fmt.Println("Your dessert is ready")
}

type classicChef struct {
	isCooking bool
}

func (c *classicChef) cook() {
	c.isCooking = true
	fmt.Println("chef started cooking your order")
	time.Sleep(time.Second * 2)
	fmt.Println("Your order is ready")
}

func main() {
	classicChef := &classicChef{}
	dessertChef := &dessertChef{}

	dessertOrder := &orderCommand{
		chef: dessertChef,
	}
	classicOrder := &orderCommand{
		chef: classicChef,
	}

	sweetToothClient := &client{
		command: dessertOrder,
	}
	classicClient := &client{
		command: classicOrder,
	}

	fmt.Println()
	sweetToothClient.invokeCommand()
	fmt.Println()
	classicClient.invokeCommand()
}
