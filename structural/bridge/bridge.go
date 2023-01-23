package main

import (
	"errors"
	"fmt"
)

type DrawAPI interface {
	DrawRectangle(int, int, int)
}

type RedRectangle struct {}

func (r *RedRectangle) DrawRectangle(radius int, x int, y int) {
	fmt.Printf("drawing shape Rectangle with red color, radius: %v, x: %v, y: %v\n", radius, x, y)
}

type GreenRectangle struct {}

func (g *GreenRectangle) DrawRectangle(radius int, x int, y int) {
	fmt.Printf("drawing shape Rectangle with green color, radius: %v, x: %v, y: %v\n", radius, x, y)
}

type Shape interface {
	Draw()
}

type Rectangle struct {
	x, y, radius int
	draw DrawAPI
}

func (c *Rectangle) Draw() error {
	if c.draw == nil {
		return errors.New("Drawer not initialized.")
	}
	
	c.draw.DrawRectangle(c.radius, c.x, c.y)
	return nil
}

func main() {
	red := &Rectangle{100, 100, 10, &RedRectangle{}}
	green := &Rectangle{100, 100, 10, &GreenRectangle{}}

	red.Draw()
	green.Draw()
}
