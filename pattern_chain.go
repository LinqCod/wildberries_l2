package main

import (
	"fmt"
	"time"
)

type customer struct {
	name              string
	troubleDifficulty int
}

type support interface {
	execute(*customer)
	setNext(support)
}

type supportBot struct {
	next support
}

func (s *supportBot) execute(c *customer) {
	if c.troubleDifficulty < 2 {
		fmt.Printf("Support bot solved %s's problem. Rate work from 1 to 5!", c.name)
		return
	}
	fmt.Println("Support bot can't help you. Redirecting to standard support...")
	time.Sleep(time.Second * 2)
	s.next.execute(c)
}
func (s *supportBot) setNext(n support) {
	s.next = n
}

type standardSupport struct {
	next support
}

func (s *standardSupport) execute(c *customer) {
	if c.troubleDifficulty < 5 {
		fmt.Printf("Standard support solved %s's problem. Rate work from 1 to 5!", c.name)
		return
	}
	fmt.Println("Standard support can't help you. Redirecting to professional support...")
	time.Sleep(time.Second * 2)
	s.next.execute(c)
}
func (s *standardSupport) setNext(n support) {
	s.next = n
}

type professionalSupport struct {
	next support
}

func (s *professionalSupport) execute(c *customer) {
	if c.troubleDifficulty <= 10 {
		fmt.Printf("Professional support solved %s's problem. Rate work from 1 to 5!", c.name)
		return
	}
	fmt.Println("Support can't help you, sorry...")
}
func (s *professionalSupport) setNext(n support) {
	s.next = n
}

func main() {
	professionalSupport := &professionalSupport{}

	standardSupport := &standardSupport{}
	standardSupport.setNext(professionalSupport)

	supportBot := &supportBot{}
	supportBot.setNext(standardSupport)

	customer := &customer{
		name:              "Maxim",
		troubleDifficulty: 11,
	}

	supportBot.execute(customer)

}
