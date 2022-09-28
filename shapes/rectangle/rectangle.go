package rectangle

import (
	"github.com/fatih/color"
)

type Renderable interface {
	Render()
}

type Rectangle struct {
	width  int
	height int
	color  *color.Color
}

/**
 *
 * Width, Height, Color
 *
 */
func New(params ...interface{}) *Rectangle {
	return &Rectangle{
		width:  params[0].(int),
		height: params[1].(int),
		color:  params[2].(*color.Color),
	}
}

func (r *Rectangle) Render() {
	r.color.Print("╔")
	for i := 0; i < r.width; i++ {
		r.color.Print("═")
	}
	r.color.Print("╗")
	println()
	for i := 0; i < r.height; i++ {
		r.color.Print("║")
		for i := 0; i < r.width; i++ {
			r.color.Print(" ")
		}
		r.color.Print("║")
		println()
	}
	r.color.Print("╚")
	for i := 0; i < r.width; i++ {
		r.color.Print("═")
	}
	r.color.Print("╝")
}

func (r *Rectangle) Width() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) Height() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}
