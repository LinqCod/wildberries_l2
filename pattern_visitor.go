package main

import "fmt"

type visitor interface {
	visitForHoney(*honey)
	visitForNuts(*nuts)
}

type product interface {
	getType() string
	accept(visitor)
}

type honey struct {
	price  float64
	amount int
}

func (h *honey) accept(v visitor) {
	v.visitForHoney(h)
}

type nuts struct {
	price  float64
	weight float64
}

func (n *nuts) accept(v visitor) {
	v.visitForNuts(n)
}

type discountCalculatorVisitor struct {
	discount float64
}

func (d *discountCalculatorVisitor) visitForHoney(h *honey) {
	fmt.Println("Starting calculation of discount for honey")
	if h.amount > 5 {
		d.discount = h.price * 0.25
	} else {
		d.discount = 0
	}
	fmt.Printf("End price of honey with discount: %0.2f\n", h.price-d.discount)
}

func (d *discountCalculatorVisitor) visitForNuts(n *nuts) {
	fmt.Println("Starting calculation of discount for nuts")
	if n.weight > 1.5 {
		d.discount = n.price * 0.5
	} else if n.weight > 1 {
		d.discount = n.price * 0.25
	} else {
		d.discount = 0
	}
	fmt.Printf("End price of nuts with discount: %0.2f\n", n.price-d.discount)
}

func main() {
	honey := &honey{
		price:  2500,
		amount: 6,
	}
	nuts := &nuts{
		price:  1500,
		weight: 1.2,
	}

	discountCalculator := &discountCalculatorVisitor{}

	fmt.Printf("\nStart price for honey: %0.2f\n", honey.price)
	honey.accept(discountCalculator)
	fmt.Printf("\nStart price for nuts: %0.2f\n", nuts.price)
	nuts.accept(discountCalculator)
}
