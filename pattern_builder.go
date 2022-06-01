package main

import "fmt"

type iBuilder interface {
	setBody()
	setGraphics()
	setProcessor()
	setRAM()
	setHardDrive()
	getComputer() computer
}

func getBuilder(builderType string) iBuilder {
	switch builderType {
	case "graphics":
		return newGraphicsBuilder()
	case "gaming":
		return newGamingBuilder()
	default:
		return nil
	}
}

type graphicsBuilder struct {
	body      string
	graphics  string
	processor string
	ram       int
	hardDrive int
}

func newGraphicsBuilder() *graphicsBuilder {
	return &graphicsBuilder{}
}

func (b *graphicsBuilder) setBody() {
	b.body = "basic body"
}
func (b *graphicsBuilder) setGraphics() {
	b.graphics = "RTX 3080ti"
}
func (b *graphicsBuilder) setProcessor() {
	b.processor = "intel core i7 8700k"
}
func (b *graphicsBuilder) setRAM() {
	b.ram = 32
}
func (b *graphicsBuilder) setHardDrive() {
	b.hardDrive = 2048
}
func (b *graphicsBuilder) getComputer() computer {
	return computer{
		body:      b.body,
		graphics:  b.graphics,
		processor: b.processor,
		ram:       b.ram,
		hardDrive: b.hardDrive,
	}
}

type gamingBuilder struct {
	body      string
	graphics  string
	processor string
	ram       int
	hardDrive int
}

func newGamingBuilder() *gamingBuilder {
	return &gamingBuilder{}
}

func (b *gamingBuilder) setBody() {
	b.body = "LED body"
}
func (b *gamingBuilder) setGraphics() {
	b.graphics = "RTX 3080ti"
}
func (b *gamingBuilder) setProcessor() {
	b.processor = "intel core i9 9900k"
}
func (b *gamingBuilder) setRAM() {
	b.ram = 64
}
func (b *gamingBuilder) setHardDrive() {
	b.hardDrive = 2048
}
func (b *gamingBuilder) getComputer() computer {
	return computer{
		body:      b.body,
		graphics:  b.graphics,
		processor: b.processor,
		ram:       b.ram,
		hardDrive: b.hardDrive,
	}
}

type computer struct {
	body      string
	graphics  string
	processor string
	ram       int
	hardDrive int
}

func (c *computer) getInfo() {
	fmt.Printf(" - body: %s\n", c.body)
	fmt.Printf(" - graphics: %s\n", c.graphics)
	fmt.Printf(" - processor: %s\n", c.processor)
	fmt.Printf(" - RAM: %d\n", c.ram)
	fmt.Printf(" - hardDrive: %d\n", c.hardDrive)

}

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildComputer() computer {
	d.builder.setBody()
	d.builder.setGraphics()
	d.builder.setProcessor()
	d.builder.setRAM()
	d.builder.setHardDrive()
	return d.builder.getComputer()
}

func main() {
	graphicsBuilder := getBuilder("graphics")
	gamingBuilder := getBuilder("gaming")

	director := newDirector(graphicsBuilder)
	graphicsComputer := director.buildComputer()

	fmt.Println("\nGraphics Computer:")
	graphicsComputer.getInfo()

	director.setBuilder(gamingBuilder)
	gamingComputer := director.buildComputer()

	fmt.Println("\nGaming Computer:")
	gamingComputer.getInfo()
}
