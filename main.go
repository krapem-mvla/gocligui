package main

import (
	"gocligui/render"
	shapes "gocligui/shapes/rectangle"

	"github.com/fatih/color"
)

func main() {
	r := render.New()

	box := shapes.New(10, 10, color.New(color.FgRed))

	r.AddChild(*box)

	r.Render()

}
